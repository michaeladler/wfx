# Secured local deployment

Local test stack:

- nginx terminates TLS on port `8443`
- Authelia provides OIDC at `https://authelia.localhost:8443`, with Redis-backed sessions
- wfx serves embedded UI at `/ui/` and uses Authelia directly as OIDC provider
- UI bearer tokens are authorized by oauth2-proxy, then forwarded to northbound port `8081`
- southbound client bearer tokens are authorized by oauth2-proxy, then forwarded to port `8080`
- UI user has no client ID claim, so northbound requests can see all jobs
- southbound client user has a `https://siemens.github.io/wfx/client-id` claim
- nginx replaces `X-Client-Id` with that claim before proxying southbound API requests, limiting access to that client's jobs
- unauthenticated API calls get `401`; UI starts OIDC authorization-code flow with PKCE
- wfx and supporting services listen only on loopback behind nginx

## Start

Start profile:

```sh
devenv --profile secured up
```

`minica` generates local CA, TLS certificate, and OIDC signing certificate under
`.devenv/profiles/secured/state/secured-deployment/certs`. Launch an isolated Firefox profile
which trusts this CA after stack becomes ready:

```sh
devenv --profile secured shell secured-browser
```

Alternatively, trust `minica.pem` in another browser or pass it to clients with `--cacert`.
`*.localhost` names resolve to loopback on common systems; otherwise add `wfx-mgmt.localhost`,
`wfx.localhost`, and `authelia.localhost` to `/etc/hosts`.

Open <https://wfx-mgmt.localhost:8443/ui/>. Generated demo credentials and OIDC client details are in
`.devenv/profiles/secured/state/secured-deployment/credentials.env`.

```sh
cat .devenv/profiles/secured/state/secured-deployment/credentials.env
```

Sign into `/ui/` with `WFX_UI_USERNAME` and `WFX_UI_PASSWORD`. This user's token has no
`https://siemens.github.io/wfx/client-id` claim, so UI can see all jobs. UI client ID remains `wfx-ui`;
its scopes are `openid profile email groups offline_access write:jobs`.

Use `WFX_CLIENT_USERNAME` and `WFX_CLIENT_PASSWORD` for southbound clients. This user's token contains:

```json
{
  "https://siemens.github.io/wfx/client-id": "e104916114ea48329ef47bf3a17fe44f"
}
```

Requests using that token on southbound API can see only jobs belonging to this client ID.

## Use an OIDC token with wfxctl

Request a token with the public `wfxctl` client using OAuth 2.0 Authorization Code Grant with PKCE. `oauth2c`
prints an Authelia URL for sign-in, then returns the token response:

```bash
TOKEN="$(oauth2c https://authelia.localhost:8443 \
  --client-id wfxctl \
  --grant-type authorization_code \
  --auth-method none \
  --audience wfx-ui \
  --response-types code \
  --redirect-url http://localhost:9876/callback \
  --pkce \
  --scopes openid,profile,email,groups,write:jobs \
  --tls-root-ca .devenv/profiles/secured/state/secured-deployment/certs/minica.pem \
  --no-browser \
  --silent | jq -r .access_token)"
wfxctl \
  --host https://wfx-mgmt.localhost:8443 \
  --tls-ca .devenv/profiles/secured/state/secured-deployment/certs/minica.pem \
  --header "Authorization: Bearer $TOKEN" \
  version
unset TOKEN
```

Open the printed URL in the secured browser, then sign in with `WFX_UI_USERNAME` and `WFX_UI_PASSWORD`. Access tokens
are short-lived; rerun the command after expiry.

Northbound API: `https://wfx-mgmt.localhost:8443/api/wfx/v1/`. Southbound API:
`https://wfx.localhost:8443/api/wfx/v1/`.

Direct wfx ports `8080` and `8081`, Authelia port `9091`, oauth2-proxy port `4180`, and Redis port `6379`
bind to loopback. Port `8443` is public entry point. devenv generates secrets and passwords on first use,
then keeps them in profile state so restarts preserve logins and sessions. Delete
`.devenv/profiles/secured/state/secured-deployment` to rotate everything. This deployment is not suitable
for production.
