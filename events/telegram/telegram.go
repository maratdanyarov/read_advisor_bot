package telegram

import "read_advisor_bot/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
	// storage
}

func New(clent *telegram.Client, storage)
