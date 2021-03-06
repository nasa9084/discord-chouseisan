package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var (
	token = "Bot " + os.Getenv("BOT_TOKEN")
)

func main() {
	if err := execute(); err != nil {
		log.Fatal(err)
	}
}

func execute() error {
	discord, err := discordgo.New()
	if err != nil {
		return fmt.Errorf("creating new session: %w", err)
	}

	discord.Token = token

	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("bot is ready")
	})

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	return nil
}
