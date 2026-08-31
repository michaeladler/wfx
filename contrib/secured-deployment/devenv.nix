# SPDX-FileCopyrightText: 2026 Siemens AG
#
# SPDX-License-Identifier: Apache-2.0
#
# Author: Michael Adler <michael.adler@siemens.com>
{
  pkgs,
  config,
  ...
}:

let
  root = config.devenv.root;
  state = "${config.devenv.state}/secured-deployment";
  render = name: file:
    pkgs.writeText name (
      builtins.replaceStrings
        [ "@ROOT@" "@STATE@" ]
        [ root state ]
        (builtins.readFile file)
    );
  autheliaConfig = render "wfx-authelia.yml" ./authelia/configuration.yml;
  oauth2ProxyConfig = render "wfx-oauth2-proxy.cfg" ./oauth2-proxy.cfg;
  oauth2ProxyAlphaConfig = render "wfx-oauth2-proxy-alpha.yml" ./oauth2-proxy-alpha.yml;
in
{
  packages = [
    pkgs.authelia
    pkgs.curl
    pkgs.firefox
    pkgs.jq
    pkgs.minica
    pkgs.nginx
    pkgs.nssTools
    pkgs.oauth2-proxy
    pkgs.oauth2c
  ];

  tasks."secured:secrets".exec = ''
    set -eu
    mkdir -p "${state}"
    umask 077

    if [ ! -f "${state}/secrets.env" ]; then
      cat > "${state}/secrets.env" <<EOF
    AUTHELIA_SESSION_SECRET=$(${pkgs.openssl}/bin/openssl rand -hex 32)
    AUTHELIA_STORAGE_ENCRYPTION_KEY=$(${pkgs.openssl}/bin/openssl rand -hex 32)
    AUTHELIA_IDENTITY_VALIDATION_RESET_PASSWORD_JWT_SECRET=$(${pkgs.openssl}/bin/openssl rand -hex 32)
    AUTHELIA_IDENTITY_PROVIDERS_OIDC_HMAC_SECRET=$(${pkgs.openssl}/bin/openssl rand -hex 32)
    OAUTH2_PROXY_COOKIE_SECRET=$(${pkgs.openssl}/bin/openssl rand -base64 32 | tr '+/' '-_')
    EOF
    fi

    if [ ! -f "${state}/credentials.env" ] \
      || [ ! -f "${state}/authelia-users.yml" ] \
      || ! ${pkgs.gnugrep}/bin/grep -q 'wfx-ui:' "${state}/authelia-users.yml" \
      || [ ! -f "${state}/oauth2-client-secret" ] \
      || [ ! -f "${state}/oauth2-client-secret-digest" ]; then
      credentials="$(${pkgs.coreutils}/bin/mktemp -d "${state}/credentials.XXXXXX")"
      trap 'rm -rf "$credentials"' EXIT
      ui_credentials="$(${pkgs.authelia}/bin/authelia crypto hash generate argon2 --random --random.length 32)"
      ui_password="$(printf '%s\n' "$ui_credentials" | ${pkgs.gnused}/bin/sed -n 's/^Random Password: //p')"
      ui_password_digest="$(printf '%s\n' "$ui_credentials" | ${pkgs.gnused}/bin/sed -n 's/^Digest: //p')"
      southbound_credentials="$(${pkgs.authelia}/bin/authelia crypto hash generate argon2 --random --random.length 32)"
      southbound_password="$(printf '%s\n' "$southbound_credentials" | ${pkgs.gnused}/bin/sed -n 's/^Random Password: //p')"
      southbound_password_digest="$(printf '%s\n' "$southbound_credentials" | ${pkgs.gnused}/bin/sed -n 's/^Digest: //p')"
      client_credentials="$(${pkgs.authelia}/bin/authelia crypto hash generate pbkdf2 --random --random.length 32)"
      client_secret="$(printf '%s\n' "$client_credentials" | ${pkgs.gnused}/bin/sed -n 's/^Random Password: //p')"
      client_secret_digest="$(printf '%s\n' "$client_credentials" | ${pkgs.gnused}/bin/sed -n 's/^Digest: //p')"
      test -n "$ui_password" && test -n "$ui_password_digest"
      test -n "$southbound_password" && test -n "$southbound_password_digest"
      test -n "$client_secret" && test -n "$client_secret_digest"

      cat > "$credentials/credentials.env" <<EOF
    WFX_UI_USERNAME=wfx-ui
    WFX_UI_PASSWORD=$ui_password
    WFX_CLIENT_USERNAME=wfx-client
    WFX_CLIENT_PASSWORD=$southbound_password
    OAUTH2_PROXY_CLIENT_ID=wfx-proxy
    OAUTH2_PROXY_CLIENT_SECRET=$client_secret
    EOF
      cat > "$credentials/authelia-users.yml" <<EOF
    ---
    users:
      wfx-ui:
        disabled: false
        displayname: WFX UI
        password: $ui_password_digest
        email: wfx-ui@local
        groups:
          - wfx
      wfx-client:
        disabled: false
        displayname: WFX Client
        password: $southbound_password_digest
        email: wfx-client@local
        groups:
          - wfx
        extra:
          wfx_client_id: e104916114ea48329ef47bf3a17fe44f
    EOF
      printf '%s' "$client_secret" > "$credentials/oauth2-client-secret"
      printf '%s' "$client_secret_digest" > "$credentials/oauth2-client-secret-digest"
      ${pkgs.coreutils}/bin/mv "$credentials"/* "${state}/"
      rmdir "$credentials"
      trap - EXIT
    fi

    chmod 600 \
      "${state}/secrets.env" \
      "${state}/credentials.env" \
      "${state}/authelia-users.yml" \
      "${state}/oauth2-client-secret" \
      "${state}/oauth2-client-secret-digest"
  '';

  tasks."secured:certificates".exec = ''
    set -eu
    certs="${state}/certs"
    mkdir -p "$certs"
    cd "$certs"

    if [ ! -f minica.pem ] || [ ! -f minica-key.pem ] \
      || [ ! -f wfx-mgmt.localhost/cert.pem ] || [ ! -f wfx-mgmt.localhost/key.pem ] \
      || ! ${pkgs.openssl}/bin/openssl x509 -checkhost wfx.localhost -noout -in wfx-mgmt.localhost/cert.pem >/dev/null 2>&1 \
      || [ ! -f oidc-signing.local/cert.pem ] || [ ! -f oidc-signing.local/key.pem ]; then
      rm -rf minica.pem minica-key.pem wfx-mgmt.localhost wfx.localhost authelia.localhost oidc-signing.local
      ${pkgs.minica}/bin/minica --ca-alg rsa --domains wfx-mgmt.localhost,wfx.localhost,authelia.localhost
      ${pkgs.minica}/bin/minica --ca-alg rsa --domains oidc-signing.local
    fi
  '';

  tasks."secured:ui".exec = ''
    set -eu
    cd "${root}/ui"
    ${pkgs.nodejs}/bin/npm ci
    ${pkgs.gleam}/bin/gleam deps download
    ${pkgs.gleam}/bin/gleam run -m lustre/dev build --minify
    cd "${root}"
    ${pkgs.go}/bin/go generate -tags ui ./ui
  '';

  scripts.secured-browser.exec = ''
    set -eu
    profile="${state}/firefox"
    mkdir -p "$profile"
    ${pkgs.nssTools}/bin/certutil -N -d "sql:$profile" --empty-password 2>/dev/null || true
    ${pkgs.nssTools}/bin/certutil -D -d "sql:$profile" -n "wfx local CA" 2>/dev/null || true
    ${pkgs.nssTools}/bin/certutil -A -d "sql:$profile" -n "wfx local CA" -t "C,," \
      -i "${state}/certs/minica.pem"
    exec ${pkgs.firefox}/bin/firefox --no-remote --profile "$profile" \
      https://wfx-mgmt.localhost:8443/ui/
  '';

  processes.wfx-secured = {
    exec = ''
      mkdir -p "${state}"
      cd "${root}"
      exec ${pkgs.go}/bin/go run -tags ui ./cmd/wfx \
        --client-host=127.0.0.1 \
        --client-port=8080 \
        --mgmt-host=127.0.0.1 \
        --mgmt-port=8081 \
        --oauth-issuer=https://authelia.localhost:8443 \
        --oauth-client-id=wfx-ui \
        --oauth-scope="openid profile email groups offline_access write:jobs" \
        --storage=sqlite \
        --storage-opt="file:${state}/wfx.db?_fk=1&_journal=WAL"
    '';
    ready.http.get = {
      port = 8081;
      path = "/health";
    };
    after = [
      "secured:ui"
      "devenv:processes:oauth2-proxy"
    ];
  };

  processes.authelia = {
    exec = ''
      set -a
      . "${state}/secrets.env"
      set +a
      exec ${pkgs.authelia}/bin/authelia \
        --config.experimental.filters=template \
        --config=${autheliaConfig}
    '';
    after = [
      "secured:certificates"
      "secured:secrets"
      "devenv:processes:redis"
    ];
    ready.http.get = {
      port = 9091;
      path = "/api/health";
    };
  };

  services.redis = {
    enable = true;
    bind = "127.0.0.1";
    port = 6379;
  };

  services.nginx = {
    enable = true;
    httpConfig = builtins.replaceStrings
      [ "@STATE@" ]
      [ state ]
      (builtins.readFile ./nginx.conf);
  };
  processes.nginx = {
    after = [
      "secured:certificates"
      "devenv:processes:authelia"
    ];
    ready.exec = ''
      ${pkgs.curl}/bin/curl --fail --silent \
        --cacert ${state}/certs/minica.pem \
        https://authelia.localhost:8443/api/health >/dev/null
    '';
  };

  processes.oauth2-proxy = {
    exec = ''
      set -a
      . "${state}/secrets.env"
      set +a
      exec ${pkgs.oauth2-proxy}/bin/oauth2-proxy \
        --config=${oauth2ProxyConfig} \
        --alpha-config=${oauth2ProxyAlphaConfig}
    '';
    after = [
      "secured:secrets"
      "devenv:processes:nginx"
    ];
    ready.http.get = {
      port = 4180;
      path = "/ping";
    };
  };
}
