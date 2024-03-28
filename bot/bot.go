package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"murvoth/legend-bot/commands"

	"github.com/bwmarrin/discordgo"
)

var (
	BotToken string
)

func Run() {
	discord, err := discordgo.New("Bot" + BotToken)
	if err != nil {
		log.Fatalf("error creating session: %v", err)
	}

	discord.AddHandler(commands.PingHandler)

	fmt.Println("adding commands...")

	commands.InitialiseSlashCommands(discord)

	if err := discord.Open(); err != nil {
		log.Fatalf("error open connection: %v", err)
	}
	defer discord.Close()

	fmt.Println("Bot running...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}