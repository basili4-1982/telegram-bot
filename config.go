package main

import (
	"os"
)

type Config struct {
	BotToken string
}

func LoadConfig() *Config {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		panic("BOT_TOKEN environment variable not set")
	}

	return &Config{
		BotToken: token,
	}
}
