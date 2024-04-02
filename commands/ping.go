package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var PingCommand = &discordgo.ApplicationCommand {
	Name: "ping",
	Description: "Replies with Pong!",
}

func PingHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	fmt.Println("running ping command")
	if i.ApplicationCommandData().Name == "ping" {
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{Content: "Pong!"},
		})
		if err != nil {
			fmt.Println(err)
		}
	}
}