package handlers

import (
	"log"
	"strings"

	tele "gopkg.in/telebot.v3"
)

// RegisterHandlers регистрирует все обработчики
func RegisterHandlers(b *tele.Bot) {
	// Команды
	b.Handle("/start", StartHandler)
	b.Handle("/help", HelpHandler)
	b.Handle("/info", InfoHandler)

	// Обработчик текстовых сообщений
	b.Handle(tele.OnText, TextHandler)

	// Обработчик команд без слэша
	b.Handle("привет", HelloHandler)

	// Обработчик callback-кнопок
	b.Handle(tele.OnCallback, CallbackHandler)
}

// StartHandler обработчик команды /start
func StartHandler(c tele.Context) error {

	return c.Send(msg, replyMarkup)
}

// HelpHandler обработчик команды /help
func HelpHandler(c tele.Context) error {
	help := "📚 Доступные команды:\n\n" +
		"/start - начать работу\n" +
		"/help - показать это сообщение\n" +
		"/info - информация о боте\n\n" +
		"Просто напишите 'привет' или любой текст!"

	return c.Send(help)
}

// InfoHandler обработчик команды /info
func InfoHandler(c tele.Context) error {
	info := "🤖 Информация о боте:\n\n" +
		"Версия: 1.0.0\n" +
		"Язык: Go\n" +
		"Библиотека: telebot v3"

	return c.Send(info)
}

// TextHandler обработчик всех текстовых сообщений
func TextHandler(c tele.Context) error {
	text := c.Text()
	log.Printf("Received message: %s from %s", text, c.Sender().Username)

	// Пример обработки разных сообщений
	switch {
	case strings.Contains(strings.ToLower(text), "как дела"):
		return c.Send("У меня всё отлично! Спасибо, что спросили 😊")

	case strings.Contains(strings.ToLower(text), "пока"):
		return c.Send("До свидания! Приходите ещё 👋")

	default:
		return c.Send("Вы написали: " + text + "\n\nИспользуйте /help для списка команд")
	}
}

// HelloHandler обработчик текстового сообщения "привет"
func HelloHandler(c tele.Context) error {
	return c.Send("И вам привет! 👋 Рад видеть!")
}

// CallbackHandler обработчик нажатий на инлайн-кнопки
func CallbackHandler(c tele.Context) error {
	callback := c.Callback()

	switch callback.Data {
	case "more_info":
		// Ответ на callback (убирает часики)
		c.Respond(&tele.CallbackResponse{
			Text: "Открываю информацию...",
		})

		return c.Send("Это пример работы инлайн-кнопок!\nВы можете добавлять любые действия по кнопкам.")
	}

	return nil
}
