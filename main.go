package main

import (
	"flag"
	"log"

	tgClient "github.com/KingDaveII/darita-food-bot/clients/telegram"
	"github.com/KingDaveII/darita-food-bot/events/telegram"
	"github.com/KingDaveII/darita-food-bot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath))

	log.Print("service started")

	consumer := telegram.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatalf("can't start consumer: %s", err.Error())
	}

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
