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
        | proxy_pass
        |
        |--> wfx-mgmt.localhost:8443 ---> wfx northbound (:8081)
        |    (UI @ /ui/, admin jobs/workflows)
        |
        +--> wfx.localhost:8443 --------> wfx southbound (:8080)
             (client jobs scoped by X-Client-Id)
```

Local test stack:

- nginx terminates TLS on port `8443`
- Authelia provides OIDC at `https://authelia.localhost:8443`, with Redis-backed sessions
- wfx serves embedded UI at `/ui/` and uses Authelia directly as OIDC provider
- `wfx-northbound-api` audience bearer tokens are authorized by oauth2-proxy, then forwarded to northbound port `8081`
- `wfx-southbound-api` audience bearer tokens from `wfx-client-1` and `wfx-client-2` are authorized by oauth2-proxy, then forwarded to southbound port `8080`
- CLI tokens use OIDC authorization-code flow with PKCE and include rotating refresh tokens
- nginx rejects tokens whose audience does not match the requested virtual host
- nginx requires `read:jobs` or `write:jobs` for southbound GET requests, and `write:jobs` for PUT requests
- Admin user has no client ID claim, so northbound requests can see all jobs
- southbound client user has a `https://siemens.github.io/wfx/client-id` claim
- nginx replaces `X-Client-Id` with that claim before proxying southbound API requests, limiting access to that client's jobs
- unauthenticated API calls get `401`; UI starts OIDC authorization-code flow with PKCE
- supporting services listen only on loopback behind nginx; local wfx listens on all interfaces so nginx container can reach it

## Start

Start profile:

```sh
devenv --profile secured up
```

Or start supporting services with Docker Compose from this directory. In another terminal, start wfx locally:

```sh
just bootstrap
docker compose up -d
go run -tags ui ../../cmd/wfx \
  --client-host=0.0.0.0 --client-port=8080 \
  --mgmt-host=0.0.0.0 --mgmt-port=8081 \
  --oauth-issuer=https://authelia.localhost:8443 \
  --oauth-client-id=wfx-ui \
  --oauth-scope="openid profile email groups offline_access read:jobs write:jobs" \
  --storage=sqlite \
  --storage-opt="file:state/wfx.db?_fk=1&_journal=WAL"
```


`bootstrap.py` generates local CA, TLS certificate, and OIDC signing certificates.

Launch an isolated Firefox profile which trusts this CA after stack becomes ready:

```sh
just browser
```

Alternatively, trust `ca.pem` in another browser or pass it to clients with `--cacert`.
`*.localhost` names resolve to loopback on common systems; otherwise add `wfx-mgmt.localhost`,
`wfx.localhost`, and `authelia.localhost` to `/etc/hosts`.

## Users

- `wfx-admin`: for wfx-ui/northbound api
- `wfx-client-1`: clientId `3eb9312f-1d86-4669-a3e7-ccc69edb1dc0`
- `wfx-client-2`: clientId `78b062b5-0f10-48d3-bfa5-6a79d44a6d8b`

## Web UI

Open <https://wfx-mgmt.localhost:8443/ui/>.
Login as `wfx-admin` user (password: `wfx-admin`).

## Endpoints

- Northbound API: `https://wfx-mgmt.localhost:8443/api/wfx/v1/`
- Southbound API: `https://wfx.localhost:8443/api/wfx/v1/`

## Example: client-scoped job visibility

Create a workflow and one job per client on the northbound API with the UI token (which has no
client ID claim, so it can create jobs for arbitrary clients), then query the southbound API with
each client token. Each client user must see only its own job.

```sh
# 1. Obtain access and refresh tokens for the admin user. Open the printed URL to sign in.
just get-wfx-ui-tokens >wfx-ui-tokens.json
WFX_UI_TOKEN="$(jq -r .access_token wfx-ui-tokens.json)"
alias wfxctl-north="wfxctl --host https://wfx-mgmt.localhost:8443 --tls-ca state/certs/ca.pem --header \"Authorization: Bearer $WFX_UI_TOKEN\""

# 2. Create a workflow.
wfxctl-north workflow create ../../workflow/dau/wfx.workflow.dau.direct.yml

# 3. Create one job for each client.
echo '{ "title": "Job for wfx-client-1" }' | wfxctl-north job create --client-id=3eb9312f-1d86-4669-a3e7-ccc69edb1dc0 --workflow=wfx.workflow.dau.direct -
echo '{ "title": "Job for wfx-client-2" }' | wfxctl-north job create --client-id=78b062b5-0f10-48d3-bfa5-6a79d44a6d8b --workflow=wfx.workflow.dau.direct -

# 4. Obtain each client's access and refresh tokens. Open the printed URL and sign in with
# username as password (e.g. wfx-client-1 / wfx-client-1). Browser launch is not required.
just get-wfx-client-1-tokens >wfx-client-1-tokens.json
just get-wfx-client-2-tokens >wfx-client-2-tokens.json
WFX_CLIENT_1_TOKEN="$(jq -r .access_token wfx-client-1-tokens.json)"
WFX_CLIENT_2_TOKEN="$(jq -r .access_token wfx-client-2-tokens.json)"

# 5. Inspect token and observe that it contains clientId; this restricts job scope internally.
jwt decode $WFX_CLIENT_1_TOKEN
# {
#   ...
#   "https://siemens.github.io/wfx/client-id": "3eb9312f-1d86-4669-a3e7-ccc69edb1dc0"
# }
```

Query southbound API with each token:

```sh
# use ui token to query all jobs (without token you'll get 403 permission denied)
wfxctl \
  --host https://wfx-mgmt.localhost:8443 \
  --tls-ca state/certs/ca.pem \
  --header "Authorization: Bearer $WFX_UI_TOKEN" \
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
78b062b5-0f10-48d3-bfa5-6a79d44a6d8b

# wfx-admin has no client ID claim, so nginx rejects its southbound request.
wfxctl \
  --host https://wfx.localhost:8443 \
  --tls-ca state/certs/ca.pem \
  --header "Authorization: Bearer $WFX_UI_TOKEN" \
  job query
# HTTP 403
```

**Note**: nginx rejects southbound requests without a client ID claim. It extracts that claim into `X-Client-Id`; wfx uses the header to filter jobs belonging to that client, which are inaccessible to other clients.

### Token Refresh

Refresh tokens rotate on every refresh. Refreshing updates the token file with new access and refresh tokens:

```sh
just refresh-tokens wfx-client-1-tokens.json
```

## Cleanup

Delete `./state` to rotate everything.
