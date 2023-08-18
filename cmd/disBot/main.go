package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	token           = "YOUR_DISCORD_BOT_TOKEN"
	channelID       = "YOUR_DISCORD_CHANNEL_ID"
	newsAPIEndpoint = "https://example.com/api/randomFrenchNews" // Replace with your actual news source
)

func main() {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection:", err)
		return
	}

	// Keep the bot running until an interrupt signal is received.
	fmt.Println("Bot is now running. Press Ctrl+C to exit.")

	go sendDailyNews(dg)

	select {}
}

func sendDailyNews(dg *discordgo.Session) {
	for {
		now := time.Now()
		next := now.Add(time.Hour * 24)
		next = time.Date(next.Year(), next.Month(), next.Day(), 6, 0, 0, 0, next.Location())
		duration := next.Sub(now)
		time.Sleep(duration)

		news, err := fetchNewsArticle()
		if err != nil {
			fmt.Println("Error fetching news:", err)
			continue
		}

		_, err = dg.ChannelMessageSend(channelID, news)
		if err != nil {
			fmt.Println("Error sending message:", err)
		}
	}
}

func fetchNewsArticle() (string, error) {
	resp, err := http.Get(newsAPIEndpoint)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var article string
	_, err = fmt.Fscanf(resp.Body, "%s", &article)
	return article, err
}
