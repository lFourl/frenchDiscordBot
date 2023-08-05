package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	// set these up in a .env
	token         = "YOUR_BOT_TOKEN_HERE"
	newsChannelID = "YOUR_NEWS_CHANNEL_ID_HERE"
)

func main() {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	dg.AddMessageCreate(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection:", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	defer dg.Close()

	go sendMorningNews(dg)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	// We can ddd any custom commands or responses here
}

func sendMorningNews(s *discordgo.Session) {
	for {
		now := time.Now()
		if now.Hour() == 8 && now.Minute() == 0 {
			_, err := s.ChannelMessageSend(newsChannelID, "Bonjour! Voici les actualités du jour : *insérer lien de l'article*")
			if err != nil {
				fmt.Println("Error sending message:", err)
			}
		}

		time.Sleep(1 * time.Minute)
	}
}
