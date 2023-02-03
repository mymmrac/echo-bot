package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

const envPrefix = "ECHO_BOT_"

func main() {
	botToken := env("TOKEN")
	bot, err := telego.NewBot(botToken, telego.WithHealthCheck(), telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	done := make(chan struct{}, 1)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	updates, err := bot.UpdatesViaWebhook("/bot"+bot.Token(), telego.WithWebhookSet(&telego.SetWebhookParams{
		URL: env("WEBHOOK_BASE") + "/bot" + bot.Token(),
	}))
	assert(err == nil, "Get updates", err)

	bh, err := th.NewBotHandler(bot, updates, th.WithStopTimeout(time.Second*10))
	assert(err == nil, "Setup bot handler", err)

	bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {
		chatID := tu.ID(message.Chat.ID)
		_, err = bot.CopyMessage(tu.CopyMessage(chatID, chatID, message.MessageID))
		if err != nil {
			bot.Logger().Errorf("Failed to copy message: %s", err)
		}

		bot.Logger().Debugf("Copied message with ID %d in chat %d", message.MessageID, chatID.ID)
	})

	go func() {
		<-sigs
		fmt.Println("Stopping...")

		err = bot.StopWebhook()
		if err != nil {
			fmt.Println("ERROR: Stop webhook:", err)
		}

		bh.Stop()

		done <- struct{}{}
	}()

	go bh.Start()

	go func() {
		err = bot.StartWebhook(env("LISTEN_ADDRESS"))
		assert(err == nil, "Start webhook:", err)
	}()

	fmt.Println("Handling updates...")

	<-done
	fmt.Println("Done")
}

func env(name string) string {
	value, ok := os.LookupEnv(envPrefix + name)
	assert(ok, "Environment variable "+envPrefix+name+" not found")
	return value
}

func assert(ok bool, args ...any) {
	if !ok {
		fmt.Println(append([]any{"FATAL:"}, args...)...)
		os.Exit(1)
	}
}
