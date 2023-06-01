# SPDX-FileCopyrightText: 2023 Siemens AG
#
# SPDX-License-Identifier: Apache-2.0
#
# Author: Michael Adler <michael.adler@siemens.com>

THISDIR := justfile_directory()
DOCKER := env_var_or_default("DOCKER", "docker")
NIX := "nix --experimental-features 'nix-command flakes'"

# postgres

export PGUSER := env_var_or_default("PGUSER", "wfx")
export PGPASSWORD := env_var_or_default("PGPASSWORD", "secret")
export PGHOST := env_var_or_default("PGHOST", "localhost")
export PGPORT := env_var_or_default("PGPORT", "5432")
export PGDATABASE := env_var_or_default("PGDATABASE", "wfx")
export PGSSLMODE := env_var_or_default("PGSSLMODE", "disable")

# mysql

export MYSQL_USER := env_var_or_default("MYSQL_USER", "root")
export MYSQL_PASSWORD := env_var_or_default("MYSQL_PASSWORD", "root")
export MYSQL_ROOT_PASSWORD := env_var_or_default("MYSQL_PASSWORD", "root")
export MYSQL_DATABASE := env_var_or_default("MYSQL_DATABASE", "wfx")
export MYSQL_HOST := env_var_or_default("MYSQL_HOST", "localhost")

# Update dependencies
update-deps:
    #!/usr/bin/env bash
    set -euxo pipefail
    go get -u ./...
    go mod tidy

# Build the documentation
pages:
    #!/usr/bin/env bash
    set -euo pipefail
    rm -rf public hugo/public
    cd hugo
    make clean && make -j`nproc`
    npm install postcss postcss-cli autoprefixer
    hugo --minify
    mv public ../

# Serve docs
docs-serve:
    #!/bin/sh
    cd hugo && make clean && make -j`nproc` && hugo server -D

# Lint code
lint:
    #!/usr/bin/env bash
    set -euo pipefail
    export CGO_ENABLED=0
    golangci-lint run -v --build-tags=sqlite,testing
    staticcheck -tags=sqlite,testing ./...
    reuse lint || true
    go list ./... 2>/dev/null | sed -e 's,github.com/siemens/wfx/,,' | grep -v "^generated" | sort | uniq | while read -r pkg; do
        if [[ "$pkg" == *tests* ]]; then
            continue
        fi
        file_count=$(find "$pkg" -maxdepth 1 -type f -name "*.go" | wc -l)
        test_count=$(find "$pkg" -maxdepth 1 -type f -name "*_test.go" | wc -l)
        if [[ "$file_count" -gt 0 ]] && [[ "$test_count" -eq 0 ]]; then
            echo "WARN: package $pkg has $file_count file(s) but no tests"
        fi
    done

# Format code
format:
    #!/usr/bin/env bash
    set -eux
    gofumpt -l -w .
    prettier -l -w .
    just --fmt --unstable

# Generate code
generate:
    #!/usr/bin/env bash
    set -eu
    go generate ./...
    go mod tidy

# Start PostgreSQL container
postgres-start:
    #!/usr/bin/env bash
    count=`{{ DOCKER }} container ls --quiet --noheading --filter name=wfx-postgres --filter "health=healthy" | wc -l`
    if [[ $count -eq 0 ]]; then
        echo "Starting PostgreSQL"
        {{ DOCKER }} run -d --rm \
            --name wfx-postgres \
            -e "POSTGRES_USER=$PGUSER" \
            -e "POSTGRES_PASSWORD=$PGPASSWORD" \
            -e "POSTGRES_DB=$PGDATABASE" \
            -p 5432:5432 \
            --health-cmd pg_isready \
            --health-interval 3s \
            --health-timeout 5s \
            --health-retries 20 \
            docker.io/library/postgres:14-alpine
    else
        echo "PostgreSQL is already running"
    fi
    while [[ $count -eq 0 ]]; do
        count=`{{ DOCKER }} container ls --quiet --noheading --filter name=wfx-postgres --filter "health=healthy" | wc -l`
        echo "Waiting for PostgreSQL to come up..."
        sleep 3
    done

# View PostgreSQL logs
@postgres-logs:
    {{ DOCKER }} logs --color -f wfx-postgres

# Enter PostgreSQL shell inside container
@postgres-shell:
    {{ DOCKER }} exec -it wfx-postgres psql -d $PGDATABASE -U $PGUSER

# Stop PostgreSQL container
postgres-stop: (_container-stop "wfx-postgres")

# Run PostgreSQL integration tests.
postgres-integration-test:
    #!/usr/bin/env bash
    set -eux
    # note: sqlite is needed for in-memory tests
    go test -tags testing,integration,postgres,sqlite -count=1 ./...

# Start wfx and connect to Postgres database
@postgres-wfx: postgres-start
    ./wfx --log-level debug \
        --storage postgres \
        --storage-opt "host=$PGHOST port=5432 database=$PGDATABASE user=$PGUSER password=$PGPASSWORD sslmode=disable"

# Generate schema definitions for PostgreSQL
postgres-generate-schema name:
    go run -mod=mod generated/ent/migrate/main.go postgres "{{ name }}"

# Stop a container with the given name.
_container-stop name:
    #!/usr/bin/env bash
    count=`{{ DOCKER }} ps --filter name={{ name }} --quiet --noheading | wc -l`
    if [[ $count -gt 0 ]]; then
        echo "Stopping {{ name }}"
        {{ DOCKER }} stop "{{ name }}" 1>/dev/null 2>/dev/null
        {{ DOCKER }} rm "{{ name }}" 2>/dev/null || true
    else
        echo "{{ name }} is already stopped"
    fi

# Start MySQL container
mysql-start:
    #!/usr/bin/env bash
    count=`{{ DOCKER }} container ls --quiet --noheading --filter name=wfx-mysql --filter "health=healthy" | wc -l`
    if [[ $count -eq 0 ]]; then
        echo "Starting MySQL"
        {{ DOCKER }} run -d --rm \
            --name wfx-mysql \
            -e MYSQL_DATABASE \
            -e MYSQL_ROOT_PASSWORD \
            -p 3306:3306 \
            --health-cmd 'mysqladmin ping --silent' \
            --health-interval 3s \
            --health-timeout 5s \
            --health-retries 20 \
            docker.io/library/mysql:8-debian
    else
        echo "MySQL is already running"
    fi
    while [[ $count -eq 0 ]]; do
        count=`{{ DOCKER }} container ls --quiet --noheading --filter name=wfx-mysql --filter "health=healthy" | wc -l`
        echo "Waiting for MySQL to come up..."
        sleep 3
    done

# Stop MySQL container
mysql-stop: (_container-stop "wfx-mysql")

# View MySQL logs
@mysql-logs:
    {{ DOCKER }} logs --color -f wfx-mysql

# Enter MySQL shell
mysql-shell:
    {{ DOCKER }} exec -it wfx-mysql mysql -u$MYSQL_USER -p$MYSQL_PASSWORD -D $MYSQL_DATABASE

# Generate schema definitions for MySQL
mysql-generate-schema name:
    go run -mod=mod generated/ent/migrate/main.go mysql "{{ name }}"

# Run MySQL integration tests.
mysql-integration-test:
    #!/usr/bin/env bash
    set -eux
    # note: sqlite is needed for in-memory tests
    go test -tags testing,integration,mysql,sqlite -count=1 ./...

# Start wfx and connect to MySQL container.
mysql-wfx: mysql-start
    ./wfx --log-level debug \
        --storage mysql \
        --storage-opt "$MYSQL_USER:$MYSQL_PASSWORD@/$MYSQL_DATABASE"

# Generate schema definitions for SQLite
sqlite-generate-schema name:
    go run -mod=mod generated/ent/migrate/main.go sqlite "{{ name }}"

# Check links used in Markdown files.
check-md-links:
    git ls-files "*.md" | xargs -n1 markdown-link-check --config .markdown-link-check.json

# vim: ts=4 sw=4 expandtab
