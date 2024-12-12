package main

import (
	"flag"
	"log"

	"kemov/LinkKeeperBot/consumer/eventConsumer"
	"kemov/LinkKeeperBot/events/telegram"
	"kemov/LinkKeeperBot/storage/files"

	tgClient "kemov/LinkKeeperBot/clients/telegram"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)
// help - Что я могу сделать?  
// .



// 7657510414:AAFnS8a5UXpAT2WlONLe7HkIEOaCK5cfpo0
func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("Сервис запущен")

	consumer := eventConsumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("Сервис остановлен")
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"Токен для доступа к телеграм-боту",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("Токен не указан")
	}

	return *token
}
