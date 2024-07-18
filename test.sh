#!/usr/bin/env bash

# Remove the existing go.sum file
rm ./go.sum >/dev/null 2>&1

GOPRIVATE=github.com/DIN-center go get github.com/cds-amal/go-sample
GOPRIVATE=github.com/DIN-center go get github.com/DIN-center/din-sc
go run .
