package telegram

import (
	"log"
	"os"
	"owm-telegram/owm"
	"regexp"

	bottl "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Define a few constants and variable to handle different commands
const startCommand string = "/start"
const helpCommand string = "/help"
const weatherCommand string = "/isitsunny"

const botTag string = "@GrassIsGreenBot"

// Telegram Token
const telegramTokenEnv string = "TELEGRAM_BOT_TOKEN"

// weather variables responses
var reponseWeather string

// Markdown formating for answers
const Markdown = bottl.ModeMarkdown

var callbackConfing = &bottl.CallbackConfig{}

// RunBot will run permanetly using different channels to acces lobbies
func RunBot() {

	// Get the session stared with the Token
	bot, err := bottl.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	updateConfig := bottl.NewUpdate(0)
	updateConfig.Timeout = 60
	updates, err := bot.GetUpdatesChan(updateConfig)

	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {

		command := regexp.MustCompile("/[a-z]+").FindString(update.Message.Text)

		switch command {
		case startCommand:
			message := "Hoy ! 😊 \nYou share your location with me\nI give you the current weather where you are at\nAnyway you should check it out yourself, don't be lazy\n `/help` if needed or `/isitsunny` to get started"
			response := bottl.NewMessage(update.Message.Chat.ID, message)
			response.ParseMode = Markdown

			
			bot.Send(response)
		case helpCommand:
			message := "You just have to type `/isitsunny` and give you your location, if you don't want to, well I am pretty useless then 😢"
			response := bottl.NewMessage(update.Message.Chat.ID, message)
			response.ParseMode = Markdown

			bot.Send(response)
		case weatherCommand:

			button := []bottl.KeyboardButton{
				bottl.NewKeyboardButtonLocation("Share location? 📍"),
			}
			replyMarkup := bottl.NewReplyKeyboard(button)
			replyMarkup.OneTimeKeyboard = true

			response := bottl.NewMessage(update.Message.Chat.ID, "Please 😘")
			response.BaseChat.ReplyMarkup = replyMarkup

			bot.Send(response)
		default:
			if update.Message.Location != nil {

				w, err := owm.GetCurrent(
					update.Message.Location,
				)
				if err != nil {
					log.Fatalln(err)
				}
				message := w.BuildAnswer()

				response := bottl.NewMessage(update.Message.Chat.ID, message)
				response.ParseMode = Markdown
				bot.Send(response)

			}
		}
	}
}
