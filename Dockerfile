FROM golang:1.20-alpine AS build

WORKDIR /echo-bot

RUN go env -w CGO_ENABLED="0"

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o /bin/echo-bot .

FROM scratch AS release

COPY --from=build /bin/echo-bot /echo-bot

ENTRYPOINT ["/echo-bot"]
