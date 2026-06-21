package main

import (
	"context"
	"flag"
	"log"

	tgClient "github.com/KingDaveII/darita-food-bot/clients/telegram"
	eventconsumer "github.com/KingDaveII/darita-food-bot/consumer/event-consumer"
	"github.com/KingDaveII/darita-food-bot/events/telegram"
	"github.com/KingDaveII/darita-food-bot/storage/sqlite"
)

const (
	tgBotHost = "api.telegram.org"
	// storagePath = "storage"
	sqliteStoragePath = "data/sqlite/storage.db"
	batchSize         = 100
)

func main() {
	// s := files.New(storagePath)
	s, err := sqlite.New(sqliteStoragePath)
	if err != nil {
		log.Fatal("failed to initialize SQLite storage: ", err)
	}

	// TODO: Add a proper context with timeout or cancellation handling
	s.Init(context.TODO())

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		s)

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
