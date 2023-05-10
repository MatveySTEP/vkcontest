package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func Sendfrs(api *tgbotapi.BotAPI, chatId int64) {
	msg := tgbotapi.NewMessage(chatId, "")
	msg.Text = "Добро пожаловать в наш бот. Выберите одну из кнопок:"
	// создаем кнопки на первом слое
	btn1 := tgbotapi.NewInlineKeyboardButtonData("Нажми", "btn1")
	btn2 := tgbotapi.NewInlineKeyboardButtonData("на", "btn2")
	btn3 := tgbotapi.NewInlineKeyboardButtonData("любую", "btn3")
	btn4 := tgbotapi.NewInlineKeyboardButtonData("кнопку", "btn4")
	// создаем первый слой кнопок
	row1 := tgbotapi.NewInlineKeyboardRow(btn1, btn2)
	row2 := tgbotapi.NewInlineKeyboardRow(btn3, btn4)

	// создаем клавиатуру
	keyboard1 := tgbotapi.NewInlineKeyboardMarkup(row1, row2)

	// добавляем клавиатуру в сообщение

	msg.ReplyMarkup = keyboard1
	api.Send(msg)
}
func Sendscd(api *tgbotapi.BotAPI, update tgbotapi.Update) {

	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Выберите действие")
	btn1 := tgbotapi.NewInlineKeyboardButtonURL("Ну вот и все", "https://www.youtube.com/watch?v=7iHk3jHz0Lg")
	btn2 := tgbotapi.NewInlineKeyboardButtonData("Назад", "btn_ret")
	row := tgbotapi.NewInlineKeyboardRow(btn1, btn2)
	keyboard := tgbotapi.NewInlineKeyboardMarkup(row)

	msg.ReplyMarkup = keyboard
	api.Send(msg)

}
func main() {
	bot, err := tgbotapi.NewBotAPI("6043187271:AAHpFLWxNG4g8CGPbM7daJte9uhp7Dk_rLQ")
	if err != nil {
		panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	for update := range updates {

		if update.CallbackQuery != nil {
			if update.CallbackQuery.Data == "btn_ret" {
				Sendfrs(bot, update.CallbackQuery.Message.Chat.ID)
			} else {
				Sendscd(bot, update)
			}

			continue
		} else if update.Message == nil { // игнорируем неподдерживаемые типы обновлений
			continue
		}

		Sendfrs(bot, update.Message.Chat.ID)
	}
}
