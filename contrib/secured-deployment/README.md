# Secured local deployment

**Warning**: This deployment is not suitable for production. However, it serves as a PoC for how to securely deploy wfx.

## Architecture

```text
       Browser / wfxctl
       (HTTPS :8443)
             |
             v
    +-------------------------------------------------------+
    |                     nginx (:8443)   <-----------------+
    |  - terminates TLS                                     |
    |  - checks token audience per vhost via auth_request   |
    |  - enforces scope (GET: read|write, PUT: write)       |
    |  - injects X-Client-Id header for southbound jobs     |
    +---+-------------------+-------------------+-----------+
        |                   |                   |           ^
        | :9091             | /oauth2/auth      |           |
        v                   v (:4180)           |           | OIDC discovery
  +------------+     +--------------+           |           | & JWKS via
  |  Authelia  |     | oauth2-proxy |-----------+-----------+ https://authelia.localhost:8443
  |  (OIDC)    |     | (validates   |           |
  +-----+------+     |  JWT/session)|           |
        |            +--------------+           |
        v :6379                                 |
  +------------+                                |
  |   Redis    |                                |
  | (sessions) |                                |
  +------------+                                |
                                                |
        +---------------------------------------+
        | route
        |
        |--> wfx-mgmt.localhost:8443 /ui/ ---> UI_DIST directory
        |    (static UI, out-of-tree)
        |
        |--> wfx-mgmt.localhost:8443 --------> wfx northbound (:8081)
        |    (admin jobs/workflows)
        |
        +--> wfx.localhost:8443 -------------> wfx southbound (:8080)
             (client jobs scoped by X-Client-Id)
```

Local test stack:

- nginx terminates TLS on port `8443` and directly serves the static UI from `/ui/`
- Authelia provides OIDC at `https://authelia.localhost:8443`, with Redis-backed sessions
- wfx runs without UI and serves northbound (`:8081`) and southbound (`:8080`) APIs
- `wfx-northbound-api` audience bearer tokens are authorized by oauth2-proxy, then forwarded to northbound port `8081`
- `wfx-southbound-api` audience bearer tokens from `wfx-client-1` and `wfx-client-2` are authorized by oauth2-proxy, then forwarded to southbound port `8080`
- CLI and devices use OAuth 2.0 client_credentials flow directly without browser
- nginx rejects tokens whose audience does not match the requested virtual host
- nginx requires `read:jobs` or `write:jobs` for southbound GET requests, and `write:jobs` for PUT requests
- Admin client (`wfxctl`) has no client ID injected into `X-Client-Id` on northbound requests
- southbound client devices have client IDs (`3eb9312f-1d86-4669-a3e7-ccc69edb1dc0`, etc.)
- nginx sets `X-Client-Id` from the token's `client_id` claim before proxying southbound API requests, limiting access to that client's jobs
- unauthenticated API calls get `401`; UI starts OIDC authorization-code flow with PKCE
- supporting services listen only on loopback behind nginx; local wfx listens on all interfaces so nginx container can reach it

## Start

Start profile:

```sh
devenv --profile secured up
```

Or start the stack with Docker Compose:

```sh
just bootstrap
just build-ui
docker compose up -d
```


`bootstrap.py` generates local CA, TLS certificate, and OIDC signing certificates.

Launch an isolated Firefox profile which trusts this CA after stack becomes ready:

```sh
just browser
```

Alternatively, trust `ca.pem` in another browser or pass it to clients with `--cacert`.
`*.localhost` names resolve to loopback on common systems; otherwise add `wfx-mgmt.localhost`,
`wfx.localhost`, and `authelia.localhost` to `/etc/hosts`.

## Clients and Users

- `wfx-admin`: web user for wfx-ui
- `wfxctl`: admin CLI client (`client_credentials` grant) for northbound API
- `3eb9312f-1d86-4669-a3e7-ccc69edb1dc0`: device client (`wfx-client-1`, `client_credentials` grant)
- `78b062b5-0f10-48d3-bfa5-6a79d44a6d8b`: device client (`wfx-client-2`, `client_credentials` grant)

## Web UI

Open <https://wfx-mgmt.localhost:8443/ui/>.
Login as `wfx-admin` user (password: `wfx-admin`).

## Endpoints

- Northbound API: `https://wfx-mgmt.localhost:8443/api/wfx/v1/`
- Southbound API: `https://wfx.localhost:8443/api/wfx/v1/`

## Example: client-scoped job visibility

Create a workflow and one job per client on the northbound API with the admin token (which has no
client ID constraint on northbound), then query the southbound API with
each client token. Each client user must see only its own job.

```sh
# 1. Obtain access token for the admin client (pure browserless via client_credentials).
just get-wfx-admin-tokens >wfx-admin-tokens.json
WFX_ADMIN_TOKEN="$(jq -r .access_token wfx-admin-tokens.json)"
alias wfxctl-north="wfxctl --host https://wfx-mgmt.localhost:8443 --tls-ca state/certs/ca.pem --header \"Authorization: Bearer $WFX_ADMIN_TOKEN\""

# 2. Create a workflow.
wfxctl-north workflow create ../../workflow/dau/wfx.workflow.dau.direct.yml

# 3. Create one job for each client.
echo '{ "title": "Job for wfx-client-1" }' | wfxctl-north job create --client-id=3eb9312f-1d86-4669-a3e7-ccc69edb1dc0 --workflow=wfx.workflow.dau.direct -
echo '{ "title": "Job for wfx-client-2" }' | wfxctl-north job create --client-id=78b062b5-0f10-48d3-bfa5-6a79d44a6d8b --workflow=wfx.workflow.dau.direct -

# 4. Obtain each client device's access token directly without browser.
just get-wfx-client-1-tokens >wfx-client-1-tokens.json
just get-wfx-client-2-tokens >wfx-client-2-tokens.json
WFX_CLIENT_1_TOKEN="$(jq -r .access_token wfx-client-1-tokens.json)"
WFX_CLIENT_2_TOKEN="$(jq -r .access_token wfx-client-2-tokens.json)"

# 5. Inspect token and observe that client_id matches the device ID.
jwt decode $WFX_CLIENT_1_TOKEN
# {
#   ...
#   "client_id": "3eb9312f-1d86-4669-a3e7-ccc69edb1dc0"
# }
```

Query southbound API with each token:

```sh
# use admin token to query all jobs on northbound
wfxctl \
  --host https://wfx-mgmt.localhost:8443 \
  --tls-ca state/certs/ca.pem \
  --header "Authorization: Bearer $WFX_ADMIN_TOKEN" \
  --filter '.content[].clientId' \
  --raw \
  job query

# wfx-client-1 sees only its job.
wfxctl \
  --host https://wfx.localhost:8443 \
  --tls-ca state/certs/ca.pem \
  --header "Authorization: Bearer $WFX_CLIENT_1_TOKEN" \
  --filter '.content[].clientId' \
  --raw \
  job query
# 3eb9312f-1d86-4669-a3e7-ccc69edb1dc0

# wfx-client-2 sees only its job.
wfxctl \
  --host https://wfx.localhost:8443 \
  --tls-ca state/certs/ca.pem \
  --header "Authorization: Bearer $WFX_CLIENT_2_TOKEN" \
  --filter '.content[].clientId' \
  --raw \
  job query
# 78b062b5-0f10-48d3-bfa5-6a79d44a6d8b

# wfxctl admin token is rejected by southbound (audience mismatch: wfx-northbound-api).
wfxctl \
  --host https://wfx.localhost:8443 \
  --tls-ca state/certs/ca.pem \
  --header "Authorization: Bearer $WFX_ADMIN_TOKEN" \
  job query
# HTTP 403
```

**Note**: nginx rejects southbound requests without a client ID in the token. It extracts `client_id` into `X-Client-Id`; wfx uses the header to filter jobs belonging to that client, which are inaccessible to other clients.

## Cleanup

Delete `./state` to rotate everything.
