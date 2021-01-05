#!/bin/bash
set -ex

docker build -t danielmschmidt/hello-wasm:latest .
docker push danielmschmidt/hello-wasm:latest
