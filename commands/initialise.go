package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func InitialiseSlashCommands(bot *discordgo.Session, appId, guildId string) {
	registeredCommands := make([]*discordgo.ApplicationCommand, 0)

	registeredCommands = append(registeredCommands, PingCommand)
	registeredCommands = append(registeredCommands, ExpressionRoll)

	for _, i := range registeredCommands {
		fmt.Println(i.Name)
	}

	_, err := bot.ApplicationCommandBulkOverwrite(appId, guildId, registeredCommands)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("commands intialised")
}