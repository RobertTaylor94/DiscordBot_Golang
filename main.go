package main

import (
	"flag"

	"murvoth/legend-bot/bot"
)

var (
	GuildID = flag.String("guild", "", "Test Guild ID. Leaving empty will register commands globally.")
	BotToken = flag.String("token", "", "Bot access token")
)

func main() {
	flag.Parse()
	bot.Run(*GuildID, *BotToken)
}