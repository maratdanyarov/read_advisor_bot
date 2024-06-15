package main

import (
	"context"
	"flag"
	"log"
	tgClient "read_advisor_bot/clients/telegram"
	event_consumer "read_advisor_bot/consumer/event-consumer"
	"read_advisor_bot/events/telegram"
	"read_advisor_bot/storage/sqlite"
)

const (
	thBotHost         = "api.telegram.org"
	sqliteStoragePath = "data/sqlite/storage.db"
	batchSize         = 100
)

func main() {
	// s := files.New(storagePath)
	s, err := sqlite.New(sqliteStoragePath)
	if err != nil {
		log.Fatal("can't connect to storage: ", err)
	}

	if err := s.Init(context.TODO()); err != nil {
		log.Fatal("can't init storage: ", err)
	}

	eventProcessor := telegram.New(
		tgClient.New(thBotHost, mustToken()),
		s,
	)

	log.Print("service started")

	consumer := event_consumer.New(eventProcessor, eventProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	//bot -tg-bot-token 'my token'
	token := flag.String("tg-bot-token", "", "token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
