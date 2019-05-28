CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o oc2sd
docker build -t us.gcr.io/tommyde-191320/oc2sd:13 .
docker push us.gcr.io/tommyde-191320/oc2sd:13
