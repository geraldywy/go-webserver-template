# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY . ./
RUN go mod download

RUN go build -o /docker-webserver

EXPOSE 8080

CMD [ "/docker-webserver" ]