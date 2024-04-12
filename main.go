package main

import (
	"flag"

	"murvoth/legend-bot/bot"
)

var (
	GuildID  = flag.String("guild", "", "Test Guild ID. Leaving empty will register commands globally.")
	BotToken = flag.String("token", "", "Bot access token")
	AppID    = flag.String("appid", "", "App ID")
)

func main() {
	flag.Parse()
	bot.Run(*AppID, *BotToken)
}
