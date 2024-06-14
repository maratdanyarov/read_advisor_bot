package main

import (
	"flag"
	"log"
	tgClient "read_advisor_bot/clients/telegram"
	event_consumer "read_advisor_bot/consumer/event-consumer"
	"read_advisor_bot/events/telegram"
	"read_advisor_bot/storage/files"
)

const (
	thBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {
	eventProcessor := telegram.New(
		tgClient.New(thBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventProcessor, eventProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	//bot -tg-bot-token 'my token'
	token := flag.String("token-bot-token", "", "token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
