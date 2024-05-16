// Contains commands used to update user configuration

package commands

import (
	"fmt"
	"murvoth/legend-bot/utility"

	"github.com/bwmarrin/discordgo"
)

var configOptions = []*discordgo.ApplicationCommandOption{
	{ // conf view subcommand
		Type:        discordgo.ApplicationCommandOptionSubCommand,
		Name:        "view",
		Description: "view your current configuration",
	},
	{ // conf update subcommand
		Type:        discordgo.ApplicationCommandOptionSubCommand,
		Name:        "update",
		Description: "update your configuration",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "color",
				Description: "The color of your dice",
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "cream",
						Value: "cream",
					},
					{
						Name:  "purple",
						Value: "purple",
					},
					{
						Name:  "orange",
						Value: "orange",
					},
					{
						Name:  "blue",
						Value: "blue",
					},
					{
						Name:  "green",
						Value: "green",
					},
				},
			},
		},
	},
}

var ConfigCommand = &discordgo.ApplicationCommand{
	Name:        "conf",
	Description: "Configure user settings",
	Options:     configOptions,
}

func ConfigHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name == "conf" {
		sub := i.ApplicationCommandData().Options[0].Name

		switch sub {
		case "view":
			viewConfig(s, i)
		case "update":
			updateConfig(s, i)
		}
	}
}

// Send a message with the current user configuration
func viewConfig(s *discordgo.Session, i *discordgo.InteractionCreate) {
	users, err := utility.GetUserConfig()
	if err != nil {
		fmt.Println(err)
	}
	propsDesc := ""

	userId := i.Member.User.ID
	user, ok := users[userId]
	if !ok {
		propsDesc = "No properties configured"
	} else {
		propsDesc = fmt.Sprintf("Dice Color: %s", user.Color)
	}

	embed := &discordgo.MessageEmbed{
		Title: fmt.Sprintf("%s Config", i.Member.Nick),
		Author: &discordgo.MessageEmbedAuthor{
			Name:    i.Member.Nick,
			IconURL: i.Member.AvatarURL(""),
		},
		Description: propsDesc,
	}

	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				embed,
			},
			Flags: discordgo.MessageFlagsEphemeral,
		},
	}); err != nil {
		fmt.Println(err)
	}
}

// Update the user configuration
func updateConfig(s *discordgo.Session, i *discordgo.InteractionCreate) {
	opts := i.ApplicationCommandData().Options[0]
	users, err := utility.GetUserConfig()
	if err != nil {
		fmt.Println(err)
	}

	// Gets color from drop down options when calling the command
	newColor := opts.Options[0].Value.(string)
	userId := i.Member.User.ID
	user, ok := users[userId]

	// If user not found, create a new user with the desired property
	// If the user already exists, updated their configuration with the new property
	if !ok {
		newUser := utility.Properties{Color: newColor}
		users[userId] = newUser
	} else {
		user.Color = newColor
		users[userId] = user
	}

	propsDesc := fmt.Sprintf("Dice Color: %s", newColor)

	if err := utility.SaveUserConfig(users); err != nil {
		fmt.Println(err)
	}

	embed := &discordgo.MessageEmbed{
		Title: fmt.Sprintf("%s Config", i.Member.Nick),
		Author: &discordgo.MessageEmbedAuthor{
			Name:    i.Member.Nick,
			IconURL: i.Member.AvatarURL(""),
		},
		Description: propsDesc,
	}

	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				embed,
			},
			Flags: discordgo.MessageFlagsEphemeral,
		},
	}); err != nil {
		fmt.Println(err)
	}

}
