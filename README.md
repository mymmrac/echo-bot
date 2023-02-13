# Echo Bot

Simple example of webhook Telegram bot running in docker with [Telego](https://github.com/mymmrac/telego).

The bot will just copy any message sent to it.

## Configuration

Provide environment variables (in `.evn` file for example), see [`.env.example`](.env.example) for reference.

> Note: `ECHO_BOT_LISTEN_URL` env used as health check base URL

## Run

```shell
docker container run -p "443:443" mymmrac/echo-bot:latest 
```

> Note: To use `.env` file add `--env-file .env`
