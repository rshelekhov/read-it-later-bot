package main

import (
	"flag"
	"github.com/rshelekhov/read-it-later-bot/clients/telegram"
	"log"
)

// TODO: move it to flag
const tgBotHost = "api.telegram.org"

func main() {
	tgClient := telegram.NewClient(tgBotHost, mustToken())

	// fetcher = fetcher.New(tgClient)

	// processor = processor.New(tgClient)

	// consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String("token", "", "token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is required")
	}

	return *token
}
