package main

import (
	"context"
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
	"resty.dev/v3"

	"telegram-bot/buttons"
	config2 "telegram-bot/internal/config"
	"telegram-bot/internal/db"
	"telegram-bot/internal/rate"
	"telegram-bot/internal/storage"
)

func main() {
	ctx := context.Background()

	// Загружаем конфигурацию
	config := config2.LoadConfig()

	if config.BotToken == "" {
		log.Fatal("BOT_TOKEN is not set")
	}

	connectDb, err := db.OpenDb(ctx, &config.Db)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Up(connectDb)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Migrations applied")

	defer connectDb.Close()

	storage := storage.NewStorage(connectDb)

	apiRate := rate.NewRate(resty.New())

	// Настройки бота
	settings := tele.Settings{
		Token:  config.BotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	// Создаем бота
	bot, err := tele.NewBot(settings)
	if err != nil {
		log.Fatal("Failed to create bot:", err)
	}

	// Регистрируем обработчики
	//handlers.RegisterHandlers(bot)

	buttons.RegisterHandlers(bot, storage, apiRate)

	log.Println("Bot started successfully!")

	// Запускаем бота (блокирует выполнение)
	bot.Start()
}
