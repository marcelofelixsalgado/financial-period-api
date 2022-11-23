# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY .env ./
COPY api/ ./api/
COPY configs/ ./configs/
COPY pkg/ ./pkg/

RUN go build -o /financial-period-api

EXPOSE 8081

CMD [ "/financial-period-api" ]