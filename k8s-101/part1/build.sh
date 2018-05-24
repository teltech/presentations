#!/usr/bin/env bash

imageName="us.gcr.io/kube-test-204901/hello:build-1"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o service
docker build -t hello .
docker tag hello ${imageName}
gcloud docker -- push ${imageName}
docker rmi ${imageName}
rm service