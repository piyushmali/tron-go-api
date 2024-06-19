FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY tron ./tron
COPY handlers ./handlers

RUN go build -o /tron-api

EXPOSE 8080

CMD [ "/tron-api" ]
