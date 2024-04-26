package main

import (
	"flag"
	"log"

	"murvoth/legend-bot/bot"
	"murvoth/legend-bot/utility"
)

var (
	GuildID = flag.String("guild", "", "Test Guild ID. Leaving empty will register commands globally.")
	BotToken = flag.String("token", "", "Bot access token")
	AppID = flag.String("appid", "", "App ID")
)

func init() {
	users, err := utility.LoadUserConfig()
	if err != nil {log.Fatal(err)}
	utility.Users = users
}

func main() {
	// appID := os.Getenv("APP_ID")
	// botToken := os.Getenv("BOT_TOKEN")
	flag.Parse()
	bot.Run(*AppID, *GuildID, *BotToken)
}
