# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

ENV GOPATH /go

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . ./

COPY app.env /app
#COPY /internal/adder/adder.go ./internal/adder
#RUN go get -d -v ./...
#RUN go install -v ./...
RUN go mod download

#COPY /cmd/shorter_server/main_docker.go ./

RUN go build -o shorter_server ./cmd/shorter_server/main.go

EXPOSE 5432

CMD [ "./shorter_server" ]
