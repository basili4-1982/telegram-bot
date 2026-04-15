package buttons

import (
	"context"
	"fmt"
	"log"

	tele "gopkg.in/telebot.v3"

	"telegram-bot/internal/rate"
	"telegram-bot/internal/storage"
)

type handlers struct {
	storage *storage.Storage
	apiRate *rate.Rate
}

// RegisterHandlers регистрирует все обработчики
func RegisterHandlers(b *tele.Bot, storage *storage.Storage, apiRate *rate.Rate) {
	// Обработчик команды /start (также срабатывает при первом контакте)

	h := handlers{
		storage: storage,
		apiRate: apiRate,
	}

	b.Handle("/start", h.startHandler)

	// Обработчик любого текстового сообщения (включая нажатия на кнопки)
	//b.Handle(tele.OnText, TextHandler)
	//
	//// Обработчик callback-кнопок
	//b.Handle(tele.OnCallback, CallbackHandler)
}

// startHandler - показывает меню сразу при старте
func (h handlers) startHandler(c tele.Context) error {
	log.Printf("User %s started bot", c.Sender().Username)

	rateVal, err := h.apiRate.Get("BTCUSDT")
	if err != nil {
		return fmt.Errorf("error while getting rate: %w", err)
	}

	err = h.storage.AddRate(context.Background(), c.Sender().Username, rateVal)
	if err != nil {
		return fmt.Errorf("error while adding rate: %w", err)
	}

	// Создаем красивое приветственное меню с кнопками
	markup := &tele.ReplyMarkup{
		ResizeKeyboard:  true,  // Подгоняем размер кнопок
		OneTimeKeyboard: false, // Клавиатура остается после нажатия
	}

	// Создаем кнопки
	btnStart := tele.ReplyButton{
		Text: "🚀 Старт",
	}
	btnHelp := tele.ReplyButton{
		Text: "❓ Помощь",
	}
	btnInfo := tele.ReplyButton{
		Text: "ℹ️ Информация",
	}
	btnMenu := tele.ReplyButton{
		Text: "📋 Меню",
	}

	// Располагаем кнопки рядами
	markup.ReplyKeyboard = [][]tele.ReplyButton{
		{btnStart},         // Первый ряд: кнопка Старт
		{btnHelp, btnInfo}, // Второй ряд: Помощь и Информация
		{btnMenu},          // Третий ряд: Меню
	}

	// Приветственное сообщение
	welcomeMsg := "🎉 *Добро пожаловать в бот!* 🎉\n\n" +
		"👇 *Нажмите кнопку СТАРТ чтобы начать работу* 👇\n\n" +
		"Или выберите одну из кнопок ниже:"

	return c.Send(welcomeMsg, markup, tele.ModeMarkdown)
}

// TextHandler - обрабатывает нажатия на все кнопки и текстовые сообщения
//func TextHandler(c tele.Context) error {
//	text := c.Text()
//
//	// Обработка нажатий на Reply-кнопки
//	switch text {
//	case "🚀 Старт":
//		return StartGameHandler(c)
//
//	case "❓ Помощь":
//		return HelpHandler(c)
//
//	case "ℹ️ Информация":
//		return InfoHandler(c)
//
//	case "◀️ Назад в меню":
//		return startHandler(c)
//
//	case "🎮 Играть":
//		return StartGameHandler(c)
//
//	case "📞 Контакты":
//		return ContactsHandler(c)
//
//	case "⚙️ Настройки":
//		return SettingsHandler(c)
//
//	default:
//		// Обработка обычного текста
//		return handleNormalText(c)
//	}
//}
//
//// StartGameHandler - основное меню после нажатия "Старт"
//func StartGameHandler(c tele.Context) error {
//	// Создаем инлайн-клавиатуру (под сообщением)
//	inlineMarkup := &tele.ReplyMarkup{}
//
//	// Кнопки действий
//	btnPlay := inlineMarkup.Data("🎮 Играть", "play_game")
//	btnRating := inlineMarkup.Data("🏆 Рейтинг", "show_rating")
//	btnRules := inlineMarkup.Data("📖 Правила", "show_rules")
//	btnBack := inlineMarkup.Data("◀️ Назад", "back_to_menu")
//
//	inlineMarkup.Inline(
//		inlineMarkup.Row(btnPlay),
//		inlineMarkup.Row(btnRating),
//		inlineMarkup.Row(btnRules),
//		inlineMarkup.Row(btnBack),
//	)
//
//	msg := "🚀 *Вы в главном меню!*\n\n" +
//		"Выберите действие:\n\n" +
//		"🎮 *Играть* - начать новую игру\n" +
//		"🏆 *Рейтинг* - посмотреть таблицу лидеров\n" +
//		"📖 *Правила* - ознакомиться с правилами"
//
//	return c.Send(msg, inlineMarkup, tele.ModeMarkdown)
//}
//
//// HelpHandler - помощь
//func HelpHandler(c tele.Context) error {
//	inlineMarkup := &tele.ReplyMarkup{}
//	btnBack := inlineMarkup.Data("◀️ Назад в меню", "back_to_menu")
//	inlineMarkup.Inline(inlineMarkup.Row(btnBack))
//
//	help := "❓ *Помощь*\n\n" +
//		"Как пользоваться ботом:\n\n" +
//		"1️⃣ Нажмите кнопку СТАРТ\n" +
//		"2️⃣ Выберите игру или действие\n" +
//		"3️⃣ Следуйте инструкциям\n\n" +
//		"Доступные команды:\n" +
//		"/start - показать главное меню\n" +
//		"/help - эта справка"
//
//	return c.Send(help, inlineMarkup, tele.ModeMarkdown)
//}
//
//// InfoHandler - информация о боте
//func InfoHandler(c tele.Context) error {
//	info := "ℹ️ *Информация о боте*\n\n" +
//		"Версия: 2.0.0\n" +
//		"Язык: Go\n" +
//		"Библиотека: telebot v3\n\n" +
//		"👨‍💻 Разработчик: @yourusername\n" +
//		"📅 Дата создания: 2024\n\n" +
//		"✨ *Особенности:*\n" +
//		"• Удобное меню с кнопками\n" +
//		"• Интерактивные игры\n" +
//		"• Рейтинг игроков"
//
//	inlineMarkup := &tele.ReplyMarkup{}
//	btnBack := inlineMarkup.Data("◀️ Назад", "back_to_menu")
//	inlineMarkup.Inline(inlineMarkup.Row(btnBack))
//
//	return c.Send(info, inlineMarkup, tele.ModeMarkdown)
//}
//
//// ContactsHandler - контакты
//func ContactsHandler(c tele.Context) error {
//	contacts := "📞 *Контакты*\n\n" +
//		"По всем вопросам:\n" +
//		"📧 Email: support@example.com\n" +
//		"💬 Telegram: @support_bot\n" +
//		"🌐 Сайт: example.com\n\n" +
//		"⏰ Время ответа: 10:00 - 19:00 МСК"
//
//	inlineMarkup := &tele.ReplyMarkup{}
//	btnBack := inlineMarkup.Data("◀️ Назад", "back_to_menu")
//	inlineMarkup.Inline(inlineMarkup.Row(btnBack))
//
//	return c.Send(contacts, inlineMarkup, tele.ModeMarkdown)
//}
//
//// SettingsHandler - настройки
//func SettingsHandler(c tele.Context) error {
//	inlineMarkup := &tele.ReplyMarkup{}
//
//	btnNotify := inlineMarkup.Data("🔔 Уведомления", "toggle_notify")
//	btnLang := inlineMarkup.Data("🌐 Язык", "change_lang")
//	btnBack := inlineMarkup.Data("◀️ Назад", "back_to_menu")
//
//	inlineMarkup.Inline(
//		inlineMarkup.Row(btnNotify),
//		inlineMarkup.Row(btnLang),
//		inlineMarkup.Row(btnBack),
//	)
//
//	return c.Send("⚙️ *Настройки*\n\nВыберите параметр для настройки:", inlineMarkup, tele.ModeMarkdown)
//}
//
//// CallbackHandler - обработка нажатий на инлайн-кнопки
//func CallbackHandler(c tele.Context) error {
//	callback := c.Callback()
//	defer c.Respond() // Обязательно отвечаем на callback
//
//	switch callback.Data {
//	case "play_game":
//		return startGame(c)
//
//	case "show_rating":
//		return showRating(c)
//
//	case "show_rules":
//		return showRules(c)
//
//	case "back_to_menu":
//		return startHandler(c)
//
//	case "toggle_notify":
//		return toggleNotifications(c)
//
//	case "change_lang":
//		return changeLanguage(c)
//
//	default:
//		return c.Edit("❌ Неизвестная команда")
//	}
//}
//
//// Игровые функции
//
//func startGame(c tele.Context) error {
//	// Простая игра "Угадай число"
//	markup := &tele.ReplyMarkup{}
//
//	btn1 := markup.Data("1", "guess_1")
//	btn2 := markup.Data("2", "guess_2")
//	btn3 := markup.Data("3", "guess_3")
//	btn4 := markup.Data("4", "guess_4")
//	btn5 := markup.Data("5", "guess_5")
//	btnBack := markup.Data("◀️ Выйти", "back_to_menu")
//
//	markup.Inline(
//		markup.Row(btn1, btn2, btn3),
//		markup.Row(btn4, btn5),
//		markup.Row(btnBack),
//	)
//
//	return c.Edit("🎮 *Игра: Угадай число*\n\nЯ загадал число от 1 до 5. Попробуй угадать!", markup, tele.ModeMarkdown)
//}
//
//func showRating(c tele.Context) error {
//	rating := "🏆 *Таблица лидеров*\n\n" +
//		"1️⃣ Игрок1 - 1500 очков\n" +
//		"2️⃣ Игрок2 - 1200 очков\n" +
//		"3️⃣ Игрок3 - 900 очков\n\n" +
//		"📊 Сыграйте и попадите в топ!"
//
//	markup := &tele.ReplyMarkup{}
//	btnBack := markup.Data("◀️ Назад", "back_to_menu")
//	markup.Inline(markup.Row(btnBack))
//
//	return c.Edit(rating, markup, tele.ModeMarkdown)
//}
//
//func showRules(c tele.Context) error {
//	rules := "📖 *Правила игры*\n\n" +
//		"1. Бот загадывает число от 1 до 5\n" +
//		"2. Вам нужно угадать число\n" +
//		"3. За правильный ответ +10 очков\n" +
//		"4. За неправильный -5 очков\n\n" +
//		"🎯 Удачи!"
//
//	markup := &tele.ReplyMarkup{}
//	btnBack := markup.Data("◀️ Назад", "back_to_menu")
//	markup.Inline(markup.Row(btnBack))
//
//	return c.Edit(rules, markup, tele.ModeMarkdown)
//}
//
//func toggleNotifications(c tele.Context) error {
//	return c.Edit("🔔 Настройки уведомлений будут доступны в следующей версии!")
//}
//
//func changeLanguage(c tele.Context) error {
//	return c.Edit("🌐 Смена языка будет доступна в следующей версии!")
//}
//
//func handleNormalText(c tele.Context) error {
//	text := c.Text()
//	lowerText := strings.ToLower(text)
//
//	if strings.Contains(lowerText, "привет") || strings.Contains(lowerText, "здравствуй") {
//		return c.Send("👋 Привет! Нажми кнопку 🚀 Старт чтобы начать игру!")
//	}
//
//	if strings.Contains(lowerText, "как дела") {
//		return c.Send("😊 Отлично! Готов к игре! Нажми 🚀 Старт")
//	}
//
//	return c.Send(fmt.Sprintf(
//		"❓ Я не понял команду: '%s'\n\nНажмите кнопку 🚀 Старт чтобы начать!",
//		text,
//	))
//}
//
//// ContactsHandler уже есть выше, но добавим еще один вариант
