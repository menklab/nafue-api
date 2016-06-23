FROM golang:1.6.2

WORKDIR /go/src/github.com/menkveldj/nafue-api
Add . /go/src/github.com/menkveldj/nafue-api

RUN go install github.com/menkveldj/nafue-api

EXPOSE 8080
ENTRYPOINT ["/go/bin/nafue-api"]
