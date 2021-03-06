package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var (
	token   = "Bot " + os.Getenv("BOT_TOKEN")
	guildID = os.Getenv("GUILD_ID")
)

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"hello": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionApplicationCommandResponseData{
				Content: "Hello!",
			},
		})
	},
}

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
	discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.Data.Name]; ok {
			h(s, i)
		}
	})

	if err := discord.Open(); err != nil {
		return fmt.Errorf("opening websocket connection")
	}
	defer discord.Close()

	if _, err := discord.ApplicationCommandCreate(
		discord.State.User.ID,
		guildID,
		&discordgo.ApplicationCommand{
			Name:        "hello",
			Description: "Say Hello",
		},
	); err != nil {
		return fmt.Errorf("creating application command: %w", err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	return nil
}
