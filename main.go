package main

import (
	"flag"
	tgClient "github.com/rshelekhov/read-it-later-bot/clients/telegram"
	"github.com/rshelekhov/read-it-later-bot/consumer/event_consumer"
	"github.com/rshelekhov/read-it-later-bot/events/telegram"
	"github.com/rshelekhov/read-it-later-bot/storage/fs"
	"log"
)

// TODO: move it to config
const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {

	eventProcessor := telegram.New(
		tgClient.NewClient(tgBotHost, mustToken()),
		fs.New(storagePath),
	)

	log.Print("Starting server...")

	consumer := event_consumer.New(eventProcessor, eventProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is failed: ", err)
	}
}

func mustToken() string {
	token := flag.String("token", "", "token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is required")
	}

	return *token
}
