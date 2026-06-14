package main

import (
	"flag"
	"log"

	tgClient "github.com/KingDaveII/darita-food-bot/clients/telegram"
	eventconsumer "github.com/KingDaveII/darita-food-bot/consumer/event-consumer"
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

	consumer := eventconsumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped: ", err)
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
