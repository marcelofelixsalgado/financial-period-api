# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY pkg/ ./pkg/
COPY api/ ./api/

RUN go build -o /financial-month-api

EXPOSE 8081

CMD [ "/financial-month-api" ]