FROM golang:latest AS go

WORKDIR /usr/src/app

COPY ./app/ .
RUN go build