package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"murvoth/legend-bot/commands"

	"github.com/bwmarrin/discordgo"
)

func Run(appId, guildID, token string) {
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("error creating session: %v", err)
	}

	discord.AddHandler(commands.PingHandler)
	discord.AddHandler(commands.ExpressionRollHandler)

	fmt.Println("adding commands...")

	commands.InitialiseSlashCommands(discord, appId, guildID)

	if err := discord.Open(); err != nil {
		log.Fatalf("error opening connection: %v", err)
	}

	defer discord.Close()

	fmt.Println("Bot running...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}