# syntax=docker/dockerfile:1
FROM golang:1.16-alpine3.14

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o /gotify-matrix-bot-docker

VOLUME [ "/data" ]

WORKDIR /data

CMD [ "/gotify-matrix-bot-docker" ]