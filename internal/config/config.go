package config

import (
	"os"
)

type (
	Config struct {
		BotToken string
		Db       DB
	}
	DB struct {
		DSN string
	}
)

func LoadConfig() *Config {
	c := &Config{}

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		panic("BOT_TOKEN environment variable not set")
	}
	c.BotToken = token

	dsn := os.Getenv("DSN")
	if dsn == "" {
		panic("DSN environment variable not set")
	}

	c.Db = DB{DSN: dsn}

	return c
}
