package commands

import(
	"github.com/bwmarrin/discordgo"
)

var PingCommand = &discordgo.ApplicationCommand {
	Name: "ping",
	Description: "Replies with Pong!",
}

func PingHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name == "ping" {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{Content: "Pong!"},
		})
	}
}