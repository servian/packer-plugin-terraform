#!/bin/sh

rm -rf dist/*
docker run --rm -it --privileged \
  -v $PWD:/go/src/github.com/servian/packer-plugin-terraform \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -w /go/src/github.com/servian/packer-plugin-terraform \
  -e GITHUB_TOKEN \
  goreleaser/goreleaser release
