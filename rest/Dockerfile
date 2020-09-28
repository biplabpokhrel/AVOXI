FROM golang:latest
ENV GO111MODULE=on
WORKDIR /app
COPY ./application/ .
RUN go mod download
