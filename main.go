package main

import (
	"flag"
	"fmt"
	"os"

	"murvoth/legend-bot/bot"
)

var (
	GuildID = flag.String("guild", "", "Test Guild ID. Leaving empty will register commands globally.")
	// BotToken = flag.String("token", "", "Bot access token")
	// AppID = flag.String("appid", "", "App ID")
)

func main() {
	appID := os.Getenv("APP_ID")
	botToken := os.Getenv("BOT_TOKEN")
	flag.Parse()
	fmt.Println(appID)
	fmt.Println(botToken)
	bot.Run(appID, *GuildID, botToken)
}
