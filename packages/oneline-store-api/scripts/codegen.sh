#!/usr/bin/env sh

set -eu

cd "$(dirname "$(cd "$(dirname "${BASH_SOURCE:-$0}")" && pwd)")"

swagger generate server --strict-additional-properties -f /oas/online-store.json
