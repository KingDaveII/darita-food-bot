package main

import (
	"flag"
	"log"
	"my_tele_bot/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {

	tgClient := telegram.New(tgBotHost, mustToken())

	//fetcher = fetcher.New(tgClient)

	// processor = processor.New(tgClient)

	// consumer.Start(fetcher, processor)

}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for accessing the Telegram Bot API",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
