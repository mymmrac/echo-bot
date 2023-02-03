FROM golang:1.20-alpine AS build

WORKDIR /echo-bot

RUN go env -w CGO_ENABLED="0"

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o /bin/echo-bot .

FROM alpine AS utils

RUN apk --update add ca-certificates && update-ca-certificates

FROM scratch AS release

COPY --from=utils /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /bin/echo-bot /echo-bot

ENTRYPOINT ["/echo-bot"]
