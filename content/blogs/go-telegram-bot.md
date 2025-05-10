---
title: Telegram Bot in Go
description: How to built a simple Telegram bot in Go
date: 2025-05-10
tags: [go, webdev, backend]
---

I recently built a simple Telegram bot using Go. The bot is designed to help users find information of a specific date. For example, if you send a message like "2025-01-10", the bot will respond with basic calendar information, fetched events and birthdays from Wikipedia, and a APOD image from NASA api. 

I'll walk you through how I built it, and for the sake of simplicity, I'll skip detailed implementation and just show the necessary parts. You can see the full source code here: [github.com/IndieCoderMM/tele-droids](https://github.com/IndieCoderMM/tele-droids/tree/master/chronobot)

## Approach

There are two main approaches to building a Telegram bot in Go:
1. **Using polling**: The bot continuously checks for new messages from the Telegram server.
2. **Using webhooks**: The Telegram server sends updates to your bot's URL whenever there are new messages.

I'll use the webhook approach in this example, as it's more easier to set up and deploy.

## Project Setup

1. **Create a new bot**: Start a chat with the [BotFather](https://t.me/botfather) on Telegram and create a new bot. You'll receive a token that you'll use to authenticate your bot.

2. **Set up a Go project**: Create a new directory for your bot and initialize a Go module.

```bash
mkdir my-telegram-bot
cd my-telegram-bot
go mod init my-telegram-bot
```

3. **Install *tgbotapi* **: This is a popular library for building Telegram bots in Go.

```bash
go get github.com/go-telegram-bot-api/telegram-bot-api
``` 

4. **Init new bot**: Create a new file `internal/bot/init.go` and add the following code:

```go
package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func InitBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI("DEMO_TOKEN") // replace with your bot token
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true // Set to false in production
	log.Printf("Authorized on account: %s", bot.Self.UserName)

	return bot
}
```

## Webhook Setup

Now, we'll set up the webhook. 

`SetWebhook` function requires a URL (e.g., `https://yourdomain.com/webhook`) where Telegram will send updates. 

```go
func SetWebhook(bot *tgbotapi.BotAPI, url string) error {
	_, err := bot.SetWebhook(tgbotapi.NewWebhook(url))
	if err != nil {
		return fmt.Errorf("failed to set webhook: %v", err)
	}
	return nil
}
```

Next function is `WebhookHandler` which returns a http handler for routing.

```go
func WebhookHandler(bot *tgbotapi.BotAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var update tgbotapi.Update
		// Parse incoming JSON payload (Telegram updates)
		err := json.NewDecoder(r.Body).Decode(&update)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		if update.Message == nil {
            http.Error(w, "Bad request", http.StatusBadRequest)
            return
		}

        // Handle input
		if update.Message.IsCommand() {
			handleCommand(bot, update)
		} else {
            handleText(bot, update)
		}
		w.WriteHeader(http.StatusOK)
	}
}
```

## Handlers

Here we'll define how the bot should respond to different commands and messages.

We can get the command from the message using `update.Message.Command()` and handle it accordingly. Commands are the ones starting with **/**. (e.g., /start, /help)

```go
func handleCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch update.Message.Command() {
    case "today":
        handlers.HandleToday(bot, update)
	case "start":
		handlers.HandleStart(bot, update)
	case "help":
	default:
		handlers.HandleHelp(bot, update)
	}
}
```

Below is an example to handle plain text. 

```go
func handleText(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
    input := update.Message.Text
    // Return the message text
    msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You said: "+input)
    bot.Send(msg)
}
```

Our bot will respond to `/today` command with today's date and a list of information.

```go
func HandleTodayInfo(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	t := time.Now()
	milestones := utils.DaysUntil(t) // Get the number of days until next month, year, etc.

    body := fmt.Sprintf("📅 Today: *%s* - *%s*\n", t.Format("2006-01-02"), t.Weekday().String())

	nasa, err := services.FetchNasaPhoto(t.Format("2006-01-02"))
	if err == nil {
        // Using markdown formatting
		body += fmt.Sprintf("\n🌌 NASA's Picture of the Day:\n[%s](%s)\n", nasa.Title, nasa.URL)
	} 

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, body)
	msg.ParseMode = "Markdown"
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.Send(msg)
}
```

## Main Function

This is where we put everything together. 

```go
// main.go
package main

import (
	"chronobot/internal/bot"
	"chronobot/internal/utils"
	"log"
	"net/http"
)

func main() {
	port := ":8080"
	url := "yourdomain.com" // replace with your domain

	b := bot.InitBot()

	if err := bot.SetWebhook(b, url+"/webhook"); err != nil {
		log.Fatal("Failed to set webhook:", err)
	}

	http.HandleFunc("/webhook", bot.WebhookHandler(b)) // Use the webhook handler
	log.Fatal(http.ListenAndServe(port, nil)) // Start the server
}
```

You can deploy this bot to any server that supports HTTPS. And, you can use services like [ngrok](https://ngrok.com/) to test the bot locally.

