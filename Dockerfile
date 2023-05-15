# syntax=docker/dockerfile:1

FROM golang:1.20.3

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
# COPY .env ./

COPY api/ ./api/
COPY commons/ ./commons/
COPY pkg/ ./pkg/
COPY settings/ ./settings/
COPY version/ ./version/

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /financial-period-api
RUN go build -o /financial-period-api

EXPOSE 8082

ENTRYPOINT ["/financial-period-api"]
# CMD ["tail", "-f", "/dev/null"]