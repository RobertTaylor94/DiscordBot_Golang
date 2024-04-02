package utility

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func EmbedMaker(exp, rollFor string, total int, rolls []int, user *discordgo.Member) *discordgo.MessageEmbed {

	name := user.Nick
	if rollFor != "" {
		name = fmt.Sprintf("%s for %s", user.Nick, rollFor)
	}
	rollsString := ConvertRollsToString(rolls)

	embed := &discordgo.MessageEmbed {
		Title: fmt.Sprintf("%s\nTotal: %v", exp, total),
		Author: &discordgo.MessageEmbedAuthor{
			Name: name,
			IconURL: user.AvatarURL(""),
		},
		Description: rollsString,
	}
	return embed
}