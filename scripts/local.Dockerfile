# syntax=docker/dockerfile:experimental

# This Dockerfile is meant to be used with the build_local_dep_image.sh script
# in order to build an image using the local version of dijeth

# Changes to the minimum golang version must also be replicated in
# scripts/ansible/roles/golang_base/defaults/main.yml
# scripts/build_avalanche.sh
# scripts/local.Dockerfile (here)
# Dockerfile
# README.md
# go.mod
FROM golang:1.17.4-buster

RUN mkdir -p /go/src/github.com/ava-labs

WORKDIR $GOPATH/src/github.com/ava-labs
COPY avalanchego avalanchego
COPY dijeth dijeth

WORKDIR $GOPATH/src/github.com/lasthyphen/dijetsgo
RUN ./scripts/build_avalanche.sh
RUN ./scripts/build_dijeth.sh ../dijeth $PWD/build/plugins/evm

RUN ln -sv $GOPATH/src/github.com/ava-labs/avalanche-byzantine/ /avalanchego
