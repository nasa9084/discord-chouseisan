package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	token = "Bot " + os.Getenv("BOT_TOKEN")
)

func main() {
	log.Print("start")
	defer log.Print("end")

	if err := execute(); err != nil {
		log.Fatal(err)
	}
}

func execute() error {
	discord, err := discordgo.New(token)
	if err != nil {
		return fmt.Errorf("creating new session: %w", err)
	}

	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("bot is ready")
	})

	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if !strings.HasPrefix(m.Content, fmt.Sprintf("<@!%s>", s.State.User.ID)) {
			return
		}

		log.Print("mentioned")

	})

	if err := discord.Open(); err != nil {
		return fmt.Errorf("opening websocket connection")
	}
	defer discord.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	return nil
}
