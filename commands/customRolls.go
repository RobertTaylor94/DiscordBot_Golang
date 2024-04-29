package commands

import (
	"fmt"
	"murvoth/legend-bot/utility"

	"github.com/bwmarrin/discordgo"
)

var crCommandOptions = []*discordgo.ApplicationCommandOption{
	{ // cr roll subcommand
		// Requires: name
		Type:        discordgo.ApplicationCommandOptionSubCommand,
		Name:        "roll",
		Description: "roll a custom roll",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "name",
				Description: "The name of your custom roll",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "bonus",
				Description: "Additional bonus to this roll",
			},
		},
	},
	{ // cr add subcommand
		// Requires: name, type, expression
		// Optional: dmgexpression
		Type:        discordgo.ApplicationCommandOptionSubCommand,
		Name:        "add",
		Description: "add a custom roll",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "name",
				Description: "The name of your custom roll",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "type",
				Description: "Attack roll or standard roll",
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "attack",
						Value: "attack",
					}, {
						Name:  "other",
						Value: "other",
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "expression",
				Description: "The expression of your roll eg 1d20+12",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "damageexp",
				Description: "The expression of your damage roll eg 1d6+4",
			},
		},
	},
	{ // cr delete subcommand
		Type:        discordgo.ApplicationCommandOptionSubCommand,
		Name:        "delete",
		Description: "delete a custom roll",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "roll_name",
				Description: "The name of your custom roll",
				Required:    true,
			},
		},
	},
	{ // cr check subcommand
		Type:        discordgo.ApplicationCommandOptionSubCommand,
		Name:        "check",
		Description: "check your custom rolls",
	},
}

var CustomRoll = &discordgo.ApplicationCommand{
	Name:        "cr",
	Description: "Custom rolls command group!",
	Options:     crCommandOptions,
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

func roll(s *discordgo.Session, i *discordgo.InteractionCreate) {
	customRolls, err := utility.GetCustomRolls()
	var embed *discordgo.MessageEmbed
	var file *discordgo.File
	if err != nil {
		fmt.Println(err)
	}
	rolling := i.ApplicationCommandData().Options[0].Options[0]

	bonus := ""
	if len(i.ApplicationCommandData().Options[0].Options) == 2 {
		bonus = i.ApplicationCommandData().Options[0].Options[1].StringValue()
	}
	userRolls := customRolls[i.Member.User.ID]
	found := false
	for _, r := range userRolls.Rolls {
		if r.Name == rolling.StringValue() {
			total, rolls, err := utility.ExpressionRoll(fmt.Sprintf("%s + %s", r.Expression, bonus))
			if err != nil {
				fmt.Println(err)
			}
			embed = utility.RollEmbedMaker(r.Expression, r.Name, total, rolls, i.Member)
			file, err = utility.GetDiceImage(i.Member.User.ID, rolls)
			if err != nil {
				fmt.Sprintln(err)
			}
			found = true
		}
	}

	if !found {
			if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("%s not found", rolling.StringValue()),
				},
			}); err != nil {
				fmt.Println(err)
			}
			return
	}

	if err != nil {
		fmt.Printf("error GetDiceImage(): %v", err)
		if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					embed,
				},
			},
		}); err != nil {
			fmt.Println(err)
		}
	} else {
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

func addRoll(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options[0]
	member := i.Member
	userID := i.Member.User.ID
	roll := utility.NewCR(options.Options)

	customRolls, err := utility.GetCustomRolls()
	if err != nil {
		fmt.Println(err)
	}

	user, ok := customRolls[userID]
	if !ok {
		newUserRolls := utility.UserRolls{Rolls: []utility.CustomRoll{roll}}
		customRolls[userID] = newUserRolls
	} else {
		user.Rolls = append(user.Rolls, roll)
		customRolls[userID] = user
	}

	if err := utility.SaveCustomRolls(customRolls); err != nil {
		fmt.Println(err)
	}

	embed := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("%s added", roll.Name),
		Description: fmt.Sprint(roll.Expression),
		Author: &discordgo.MessageEmbedAuthor{
			Name:    member.Nick,
			IconURL: member.AvatarURL(""),
		},
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

func deleteRoll(s *discordgo.Session, i *discordgo.InteractionCreate) {
	customRolls, err := utility.GetCustomRolls()
	if err != nil {
		fmt.Println(err)
	}
	rollToDelete := i.ApplicationCommandData().Options[0].Options[0]
	userRolls := customRolls[i.Member.User.ID]
	newRolls := utility.UserRolls{}

	for _, r := range userRolls.Rolls {
		if r.Name != rollToDelete.StringValue() {
			newRolls.Rolls = append(newRolls.Rolls, r)
		}
	}

	customRolls[i.Member.User.ID] = newRolls
	if err := utility.SaveCustomRolls(customRolls); err != nil {
		fmt.Println(err)
	}

	embed := utility.CustomDiceEmbedMaker(newRolls, i.Member)
	embed.Title = fmt.Sprintf("%s deleted", rollToDelete.StringValue())
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

func checkRolls(s *discordgo.Session, i *discordgo.InteractionCreate) {
	customRolls, err := utility.GetCustomRolls()
	if err != nil {
		fmt.Println(err)
	}
	userRolls := customRolls[i.Member.User.ID]
	embed := utility.CustomDiceEmbedMaker(userRolls, i.Member)

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
