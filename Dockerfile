# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /ALTS
EXPOSE 8080
CMD [ "/ALTS" ]