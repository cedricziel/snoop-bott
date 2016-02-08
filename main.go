package main

import (
	"os"
	"time"

	"github.com/cedricziel/snoop-bott/Godeps/_workspace/src/github.com/tucnak/telebot"
)

// main starts the program
func main() {
	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		println("No telegram bot secret token. Set one using the 'TELEGRAM_TOKEN' environment variable")
	}
	bot, err := telebot.NewBot(token)
	if err != nil {
		println(err)
		return
	}

	messages := make(chan telebot.Message)
	bot.Listen(messages, 1*time.Second)

	for message := range messages {
		if message.Text == "/hi" {
			bot.SendMessage(message.Chat,
				"Hello, "+message.Sender.FirstName+"!", nil)
		}
	}
}
