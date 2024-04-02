package commands

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var options = []*discordgo.ApplicationCommandOption{
	{
		Type: discordgo.ApplicationCommandOptionString,
		Name: "expression",
		Description: "The dice roll expression (e.g. 1d10 + 1d4 + 7)",
		Required: true,
	},
	{
		Type: discordgo.ApplicationCommandOptionString,
		Name: "roll_type",
		Description: "What are you rolling for",
		Required: false,
	},
}

var ExpressionRoll = &discordgo.ApplicationCommand{
	Name:        "q",
	Description: "Roll dice with an expression",
	Options: options,
}

func ExpressionRollHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name == "q" {
		options := i.ApplicationCommandData().Options
		expression := strings.ReplaceAll(options[0].StringValue(), " ", "")

		total := 0
		rolls := make([]int, 0)

		// split expression into component parts
		initialSplit := strings.Split(expression, "+")
		for _, v := range initialSplit {
			if strings.Contains(v, "d") {
				// split roll find type and number of dice
				roll := strings.Split(v, "d")
				num := roll[0]
				sides, err := strconv.Atoi(roll[1])
				if err != nil {fmt.Printf("Error converting sides to int: %v", err)}
				// roll dice and add value to total
				for range num {
					roll := rand.Intn(sides) + 1
					rolls = append(rolls, roll)
					total += roll
				}
			} else {
				// add bonuses to roll total
				bonus, err := strconv.Atoi(v)
				if err != nil {fmt.Println("Error converting bonus to int", err)}
				total += bonus
			}
		}

		

		if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Total: %v", total),
			},
		}); err != nil {
			fmt.Println(err)
		}
	}
}
