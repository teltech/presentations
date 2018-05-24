#!/usr/bin/env bash

imageName="us.gcr.io/kube-test-204901/srv1:build-3"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o service
docker build -t srv1 .
docker tag srv1 ${imageName}
gcloud docker -- push ${imageName}
docker rmi ${imageName}
rm service
