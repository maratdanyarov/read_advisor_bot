package main

import (
	"flag"
	"log"
	"read_advisor_bot/clients/telegram"
)

const (
	thBotHost := "api.telegram.org"
)

func main() {
	tgClient = telegram.New(mustToken())

	// fetcher = fetcher.New

	// processor = processor.New

	// consumer.Start(fetcher, processor)
}

//7464388264:AAEzgvcHZtKke3kuZW8PGZhNyv-bwhgejYI

func mustToken() string {
	//bot -tg-bot-token 'my token'
	token := flag.String("token-bot-token", "", "token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
