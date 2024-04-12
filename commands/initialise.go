package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func InitialiseSlashCommands(bot *discordgo.Session, appId string) {
	registeredCommands := make([]*discordgo.ApplicationCommand, 0)

	registeredCommands = append(registeredCommands, PingCommand)
	registeredCommands = append(registeredCommands, ExpressionRoll)
	registeredCommands = append(registeredCommands, CustomRoll)

	for _, i := range registeredCommands {
		fmt.Println(i.Name)
	}

	for _, cmd := range registeredCommands {
		_, err := bot.ApplicationCommandCreate(appId, "", cmd)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("commands intialised")
}
