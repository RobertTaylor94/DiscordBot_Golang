package utility

import (
	"github.com/bwmarrin/discordgo"
)

type CustomRoll struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Expression  string `json:"expression"`
	DamageExp   string `json:"damageExp,omitempty"`
	Bonus       int    `json:"bonus,omitempty"`
	DamageBonus string `json:"damageBonus,omitempty"`
}

func NewCR(opts []*discordgo.ApplicationCommandInteractionDataOption) CustomRoll {
	newRoll := CustomRoll{}

	for _, opts := range opts {
		switch opts.Name {
		case "name":
			newRoll.Name = opts.StringValue()
		case "type":
			newRoll.Type = opts.StringValue()
		case "expression":
			newRoll.Expression = opts.StringValue()
		case "damageexp":
			newRoll.DamageExp = opts.StringValue()
		case "bonus":
			newRoll.Bonus = int(opts.IntValue())
		case "damagebonus":
			newRoll.DamageBonus = opts.StringValue()
		}
	}
	return newRoll
}
