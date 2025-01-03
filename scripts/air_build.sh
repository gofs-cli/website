#!/usr/bin/env bash

set -e
set -u
set -o pipefail
set -x

npm run build
npm run tailwind
templ generate
CGO_ENABLED=0 go build -o ./tmp/main ./cmd/server/
