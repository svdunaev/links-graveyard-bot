package main

import (
	"context"
	"github.com/joho/godotenv"
	tgClient "links-graveyard/clients/telegram"
	eventConsumer "links-graveyard/consumer/event-consumer"
	"links-graveyard/events/telegram"
	"links-graveyard/internal/config"
	"links-graveyard/storage/sqlite"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {
	cfg := config.MustLoad()

	s, err := sqlite.New(cfg.SqliteStoragePath)
	if err != nil {
		log.Fatal("cant connect to storage: ", err)
	}

	if err := s.Init(context.Background()); err != nil {
		log.Fatal("cant init storage: ", err)
	}

	eventsProcessor := telegram.New(
		tgClient.New(cfg.Host, cfg.Token),
		s,
	)

	log.Print("service started")

	consumer := eventConsumer.New(eventsProcessor, eventsProcessor, cfg.BatchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service stopped", err)
	}
}
