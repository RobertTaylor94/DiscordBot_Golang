package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func InitialiseSlashCommands(bot *discordgo.Session, token, guildId string) {
	registeredCommands := make([]*discordgo.ApplicationCommand, 0)

	registeredCommands = append(registeredCommands, PingCommand)

	for _, i := range registeredCommands {
		fmt.Println(i.Name)
	}

	bot.ApplicationCommandBulkOverwrite(token, guildId, registeredCommands)

	fmt.Println("commands intialised")
}