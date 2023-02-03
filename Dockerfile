FROM golang:1.20-alpine AS build

RUN apk --update add ca-certificates upx && update-ca-certificates

WORKDIR /echo-bot

RUN go env -w CGO_ENABLED="0"

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -ldflags="-s -w" -o /bin/echo-bot . && upx --best --lzma /bin/echo-bot

FROM scratch AS release

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /bin/echo-bot /echo-bot

ENTRYPOINT ["/echo-bot"]
