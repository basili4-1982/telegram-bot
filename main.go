package main

import (
	"log"
	"time"

	tele "gopkg.in/telebot.v3"

	"telegram-bot/handlers"
)

func main() {
	// Загружаем конфигурацию
	config := LoadConfig()

	if config.BotToken == "" {
		log.Fatal("BOT_TOKEN is not set")
	}

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
	handlers.RegisterHandlers(bot)

	log.Println("Bot started successfully!")

	// Запускаем бота (блокирует выполнение)
	bot.Start()
}
