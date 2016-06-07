FROM golang:1.6.2

WORKDIR /go/src/nafue-api
Add . /go/src/nafue-api

RUN go get -u github.com/kardianos/govendor

RUN bash utility.sh deps

RUN go install nafue-api

EXPOSE 8080
ENTRYPOINT ["/go/bin/nafue-api"]
