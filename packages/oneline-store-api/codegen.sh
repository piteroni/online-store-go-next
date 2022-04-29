#!/usr/bin/env sh

set -eu

cd "$(dirname "$0")"

swagger generate server --strict-additional-properties -f /oas/online-store.json
