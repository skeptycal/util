#!/usr/bin/env zsh

# go build for WASM js

GOOS=js GOARCH=wasm go build -o main.wasm
