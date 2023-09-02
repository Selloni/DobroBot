package main

import (
	"DobroBot/client/telegram"
	"flag"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.New(mustToken())
}

func mustToken() string {
	token := flag.String("token",
		"6548886185:AAH_D2kYxX2GIV5PhuDWKPjwBpidWeeBVx4",
		"token for access to telegram bot")
	if *token == "" {
		log.Fatal("empty argument for flag")
	}
	return *token
}
