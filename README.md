# Echo Bot

Simple example of webhook Telegram bot running in docker with [Telego](https://github.com/mymmrac/telego).

The bot will just copy any message sent to it.

## Configuration

Provide environment variables (in `.evn` file for example):

```dotenv
ECHO_BOT_TOKEN="BOT_TOKEN"
ECHO_BOT_WEBHOOK_BASE="https://example.com"
ECHO_BOT_LISTEN_ADDRESS=":443"
```

## Run

```shell
docker container run -p "443:443" mymmrac/echo-bot:latest 
```

> Note: To use `.env` file add `--env-file .env`
