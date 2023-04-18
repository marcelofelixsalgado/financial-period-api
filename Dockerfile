# syntax=docker/dockerfile:1

FROM golang:1.19.4-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY .env ./

COPY api/ ./api/
COPY commons/ ./commons/
COPY pkg/ ./pkg/
COPY settings/ ./settings/
COPY version/ ./version/

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

RUN go build -o /financial-period-api

EXPOSE 8082

ENTRYPOINT ["financial-period-api"]