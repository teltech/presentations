FROM alpine:3.7
RUN apk update && apk add ca-certificates
COPY oc2sd /bin/.
COPY service-account.json /
CMD ["/bin/oc2sd"]