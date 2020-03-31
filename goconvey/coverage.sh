#!/bin/bash
COVERAGE_NAME="coverage"
PKG_LIST=$(go list ./... | grep -v /vendor/)

go test -covermode=count -coverpkg="${PKG_LIST}" -coverprofile "${COVERAGE_NAME}.cov"

if [ "$1" == "html" ]; then
  go tool cover -html="${COVERAGE_NAME}.cov" -o coverage.html
fi

if [ "$1" == "total" ]; then
  go tool cover -func="${COVERAGE_NAME}.cov" -o coverage.txt
fi
