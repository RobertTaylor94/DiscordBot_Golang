package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var crCommandOptions = []*discordgo.ApplicationCommandOption{
	{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "roll",
			Description: "roll a custom roll",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "roll_name",
					Description: "The name of your custom roll",
				},
			},
		},
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "add",
			Description: "add a custom roll",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "roll_name",
					Description: "The name of your custom roll",
				},
			},
		},
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "delete",
			Description: "delete a custom roll",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "roll_name",
					Description: "The name of your custom roll",
				},
			},
		},
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "check",
			Description: "check your custom rolls",
		},
}

var CustomRoll = &discordgo.ApplicationCommand{
	Name:        "cr",
	Description: "Custom rolls command group!",
	Options: crCommandOptions,
}

func CustomRollHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name == "cr" {
		subCommand := i.ApplicationCommandData().Options[0].Name

		switch subCommand {
		case "roll":
			roll(s, i)
		case "add":
			addRoll(s, i)
		case "delete":
			deleteRoll(s, i)
		case "check":
			checkRolls(s, i)
		}

	}
}

func addRoll(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: "added roll yay!"},
	}); err != nil {
		fmt.Println(err)
	}
}

func roll(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: "custom roll yay!"},
	}); err != nil {
		fmt.Println(err)
	}
}

func deleteRoll(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: "deleted custom roll yay!"},
	}); err != nil {
		fmt.Println(err)
	}
}

func checkRolls(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: "check custom rolls yay!"},
	}); err != nil {
		fmt.Println(err)
	}
}