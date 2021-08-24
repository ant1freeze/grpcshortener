FROM golang:alpine

ADD . /go/src/github.com/ant1freeze/grpcshortener/

RUN go install github.com/ant1freeze/grpcshortener/

ENTRYPOINT ["/go/bin/shorter_server"]

EXPOSE 50051
EXPOSE 5432
