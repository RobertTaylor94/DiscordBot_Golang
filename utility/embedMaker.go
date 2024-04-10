package utility

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func EmbedMaker(exp, rollFor string, total int, rolls []Roll, user *discordgo.Member) *discordgo.MessageEmbed {

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
