#syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod go.sum ./

COPY . ./

COPY ./configs/app.env /app

RUN go mod download

RUN go build -o shorter_server ./cmd/shorter_server/main.go

EXPOSE 50051

CMD ["./shorter_server"]