FROM golang:1.13-alpine as build-env
WORKDIR /cloud-native-go

RUN apk update && apk add --no-cache gcc musl-dev git bash

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app

CMD ["/cloud-native-go/bin/app"]
EXPOSE 8080