package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jessevdk/go-flags"

	tele "gopkg.in/telebot.v4"
)

type options struct {
	BotToken string `long:"token" env:"BOT_TOKEN" required:"true" description:"telegram bot token"`
}

type readBusinessConnectionMessagePayload struct {
	BusinessConnectionID string `json:"business_connection_id"`
	MessageID            int    `json:"message_id"`
	ChatID               int64  `json:"chat_id"`
}

func main() {
	var opts options
	p := flags.NewParser(&opts, flags.PrintErrors|flags.PassDoubleDash|flags.HelpFlag)
	if _, err := p.Parse(); err != nil {
		os.Exit(1)
	}

	log.Println("bot started")

	if err := run(opts); err != nil {
		log.Printf("run: %s", err)
	}

	log.Println("bot stopped")
}

func run(opts options) error {
	pref := tele.Settings{
		Token:  opts.BotToken,
		Poller: &tele.LongPoller{Timeout: time.Second * 5},
	}
	bot, err := tele.NewBot(pref)
	if err != nil {
		return fmt.Errorf("tele.NewBot: %w", err)
	}

	bot.Handle(tele.OnBusinessMessage, handle)

	bot.Start()
	return nil
}

func handle(c tele.Context) error {
	msg := c.Update().BusinessMessage
	_, err := c.Bot().Raw("readBusinessMessage", readBusinessConnectionMessagePayload{
		BusinessConnectionID: msg.BusinessConnectionID,
		MessageID:            msg.ID,
		ChatID:               msg.Sender.ID,
	})
	if err != nil {
		return fmt.Errorf("readBusinessMessage: %w", err)
	}
	return nil
}
