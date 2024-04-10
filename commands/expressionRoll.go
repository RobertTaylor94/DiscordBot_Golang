package commands

import (
	"fmt"
	"murvoth/legend-bot/utility"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var options = []*discordgo.ApplicationCommandOption{
	{
		Type:        discordgo.ApplicationCommandOptionString,
		Name:        "expression",
		Description: "The dice roll expression (e.g. 1d10 + 1d4 + 7)",
		Required:    true,
	},
	{
		Type:        discordgo.ApplicationCommandOptionString,
		Name:        "roll_type",
		Description: "What are you rolling for",
		Required:    false,
	},
}

var ExpressionRoll = &discordgo.ApplicationCommand{
	Name:        "r",
	Description: "Roll dice with an expression",
	Options:     options,
}

func ExpressionRollHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name == "r" {
		options := i.ApplicationCommandData().Options

		expression := strings.ReplaceAll(options[0].StringValue(), " ", "")

		user := i.Member
		total, rolls, err := utility.ExpressionRoll(expression)
		if err != nil {
			fmt.Println(err)
		}

		rollFor := ""
		if len(options) == 2 {
			rollFor = options[1].StringValue()
		}

		file, err := utility.GetDiceImage("", rolls)
		if err != nil {
			fmt.Printf("error GetDiceImage(): %v", err)
		}

		embed := utility.EmbedMaker(expression, rollFor, total, rolls, user)

		if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					embed,
				},
				Files: []*discordgo.File{
					file,
				},
			},
		}); err != nil {
			fmt.Println(err)
		}
	}
}
