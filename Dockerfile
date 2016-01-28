FROM golang:1.5.3

WORKDIR /go/src/nafue
Add . /go/src/nafue

RUN go get github.com/tools/godep

RUN godep restore

RUN go install nafue

EXPOSE 8080
ENTRYPOINT ["/go/bin/nafue"]
