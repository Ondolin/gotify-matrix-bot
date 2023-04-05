# syntax=docker/dockerfile:1
FROM golang:alpine AS builder

RUN apk add --no-cache git ca-certificates build-base su-exec olm-dev

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o /gotify-matrix-bot-docker

FROM alpine:latest

RUN apk add --no-cache ca-certificates olm bash

COPY --from=builder /gotify-matrix-bot-docker /gotify-matrix-bot-docker

VOLUME [ "/data" ]

WORKDIR /data

CMD [ "/gotify-matrix-bot-docker" ]
