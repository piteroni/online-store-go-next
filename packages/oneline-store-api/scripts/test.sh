#!/usr/bin/env sh

set -eu

cd "$(dirname "$0")"

[ ! -d tmp ] && mkdir tmp
gotest -v -cover ./... -coverprofile=tmp/cover.out
go tool cover -html=tmp/cover.out -o tmp/cover.html
