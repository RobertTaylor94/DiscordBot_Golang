package utility

import (

	"github.com/bwmarrin/discordgo"
)

type roll struct {
	Name        string
	Type        string
	Expression  string
	DamageExp   string
	Bonus       int
	DamageBonus string
}

func NewCR(opts []*discordgo.ApplicationCommandInteractionDataOption) roll {
	newRoll := roll{}

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