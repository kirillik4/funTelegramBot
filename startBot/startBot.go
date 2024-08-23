package startBot

import (
	"Projects/pisaBrain"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"log"
)

func StartBot(token string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}
	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Println(update.Message.From.UserName, update.Message.From.ID)
			if update.Message.Text == "/pisa" || update.Message.Text == "/pisa@JendossBot" {
				text := pisaBrain.PisaMove(update.Message.From.ID, update.Message.From.UserName)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
				bot.Send(msg)
			} else if update.Message.Text == "/toppisas" || update.Message.Text == "/toppisas@JendossBot" {
				text := pisaBrain.TopPisas()
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
				bot.Send(msg)
			}

		}
	}
}
