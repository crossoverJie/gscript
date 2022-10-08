FROM golang:1.17.13-alpine3.16 AS builder

# env
ENV PROJECT_NAME=gscript

# workdir
WORKDIR /${PROJECT_NAME}

# cache image layer
COPY go.mod .
COPY go.sum .
RUN --mount=type=ssh go mod download

COPY . .
RUN go build -ldflags "-s -w" cmd/gscript.go
RUN cp gscript /usr/bin/