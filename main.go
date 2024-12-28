package main

import (
	"context"
	"flag"
	"log"

	"kemov/LinkKeeperBot/consumer/eventConsumer"
	"kemov/LinkKeeperBot/events/telegram"
	"kemov/LinkKeeperBot/storage/sqlite"

	tgClient "kemov/LinkKeeperBot/clients/telegram"
)

const (
	tgBotHost         = "api.telegram.org"
	sqliteStoragePath = "data/sqlite/storage.db"
	batchSize         = 100
)

func main() {
	// s := files.New(storagePath)
	s, err := sqlite.New(sqliteStoragePath)
	if err != nil {
		log.Fatal("Не удается подключиться к хранилищу: ", err)
	}

	if err := s.Init(context.TODO()); err != nil {
		log.Fatal("Не удается инициализировать хранилище: ", err)
	}

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		s,
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
