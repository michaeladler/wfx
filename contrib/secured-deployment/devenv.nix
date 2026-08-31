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
  state = "${root}/contrib/secured-deployment/state";
in
{
  packages = [
    pkgs.authelia
    pkgs.curl
    pkgs.firefox
    pkgs.jq
    pkgs.nginx
    pkgs.nssTools
    pkgs.oauth2-proxy
    pkgs.oauth2c
    pkgs.nssTools
  ];

  languages.python = {
    enable = true;
    package = pkgs.python3.withPackages (ps: [
      ps.cryptography
    ]);
  };

  tasks."secured:bootstrap".exec = ''
    set -eu
    cd "${root}/contrib/secured-deployment"
    ${pkgs.python3}/bin/python3 bootstrap.py --local --state "${state}"
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
        --oauth-scope="openid profile email groups offline_access read:jobs write:jobs" \
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
        --config="${state}/authelia.yml"
    '';
    after = [
      "secured:bootstrap"
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
    httpConfig = builtins.replaceStrings [ "@STATE@" ] [ state ] (builtins.readFile ./nginx.conf);
  };
  processes.nginx = {
    after = [
      "secured:bootstrap"
      "devenv:processes:authelia"
    ];
    ready.exec = ''
      ${pkgs.curl}/bin/curl --fail --silent \
        --cacert ${state}/certs/ca.pem \
        https://authelia.localhost:8443/api/health >/dev/null
    '';
  };

  processes.oauth2-proxy = {
    exec = ''
      set -a
      . "${state}/secrets.env"
      set +a
      exec ${pkgs.oauth2-proxy}/bin/oauth2-proxy \
        --config="${state}/oauth2-proxy.cfg" \
        --alpha-config="${state}/oauth2-proxy-alpha.yml"
    '';
    after = [
      "secured:bootstrap"
      "devenv:processes:nginx"
    ];
    ready.http.get = {
      port = 4180;
      path = "/ping";
    };
  };
}
