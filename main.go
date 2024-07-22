package main

import (
	"flag"
	"log"
	// "read-adviser/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	// telegramClient := telegram.New(mustToken())
}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"Telegram access token",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("Token is not specifed.")
	}
}
