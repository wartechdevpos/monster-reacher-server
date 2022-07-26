FROM golang:1.18-alpine as builder

ARG SERVICE_NAME 

WORKDIR /app

COPY config monster-reacher-server/config
COPY microservices monster-reacher-server/microservices
COPY libraries monster-reacher-server/libraries

WORKDIR /app/monster-reacher-server

WORKDIR /app/monster-reacher-server/microservices/cmd/${SERVICE_NAME}

RUN go mod download
RUN go build .