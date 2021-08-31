# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . ./
#COPY /internal/adder/adder.go ./internal/adder
#RUN go get -d -v ./...
#RUN go install -v ./...
RUN go mod download

#COPY /cmd/shorter_server/main_docker.go ./

RUN go build -o /docker-gs-ping

EXPOSE 8080

CMD [ "/docker-gs-ping" ]
