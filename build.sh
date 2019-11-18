#!/bin/sh

go fmt
golint

export GOOS=darwin
go build -ldflags='-s -w' -o packer-provisioner-terraform
