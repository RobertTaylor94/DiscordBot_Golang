package bot

import (
	"fmt"
	"log"
	"murvoth/legend-bot/commands"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

func Run(appId, guildID, token string) {
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("error creating session: %v", err)
	}

	fmt.Println("adding commands...")
	discord.AddHandler(commands.PingHandler)
	discord.AddHandler(commands.ExpressionRollHandler)
	discord.AddHandler(commands.CustomRollHandler)
	discord.AddHandler(commands.ConfigHandler)
	commands.InitialiseSlashCommands(discord, guildID, appId)

	if err := discord.Open(); err != nil {
		log.Fatalf("error opening connection: %v", err)
	}

	defer discord.Close()

	fmt.Println("Bot running...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
