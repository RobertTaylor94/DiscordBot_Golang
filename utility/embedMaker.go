package utility

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func RollEmbedMaker(exp, rollFor string, total int, rolls []Roll, user *discordgo.Member) *discordgo.MessageEmbed {

	name := user.Nick
	if rollFor != "" {
		name = fmt.Sprintf("%s for %s", user.Nick, rollFor)
	}

	embed := &discordgo.MessageEmbed{
		Title: fmt.Sprintf("%s\n**Total: %v**", exp, total),
		Author: &discordgo.MessageEmbedAuthor{
			Name:    name,
			IconURL: user.AvatarURL(""),
		},
		Image: &discordgo.MessageEmbedImage{
			URL: "attachment://image.png",
		},
	}
	return embed
}

func CustomDiceEmbedMaker(rolls UserRolls, user *discordgo.Member) *discordgo.MessageEmbed{
	embed := &discordgo.MessageEmbed{
		Title: fmt.Sprintf("%s Custom Rolls", user.Nick),
		Author: &discordgo.MessageEmbedAuthor{
			Name:    user.Nick,
			IconURL: user.AvatarURL(""),
		},
	}
	for _, r := range rolls.Rolls {
		if r.Type == "attack" {
			emField := &discordgo.MessageEmbedField{
				Name:  r.Name,
				Value: fmt.Sprintf("Attack Bonus: %s\nDamage: %s + %s\nType: %s", r.Expression, r.DamageExp, r.DamageBonus, r.Type),
			}
			embed.Fields = append(embed.Fields, emField)
		} else {
			emField := &discordgo.MessageEmbedField{
				Name:  r.Name,
				Value: fmt.Sprintf("Roll: %s", r.Expression),
			}
			embed.Fields = append(embed.Fields, emField)
		}
	}
	return embed
}