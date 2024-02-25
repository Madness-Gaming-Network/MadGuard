package chat

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// TODO: Consider "interactions"?
// https://discord.com/developers/docs/topics/community-resources#interactions
// https://github.com/Amatsagu/Tempest

func RunBot() {
	// TODO: Fetch bot token from env var
	token := "YOUR_BOT_TOKEN"
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v\n", err)
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		log.Fatalf("Error opening Discord session: %v\n", err)
	}

	log.Printf("Bot is now running. Press CTRL-C to exit\n")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	err = dg.Close()
	if err != nil {
		log.Fatalf("Error closing Discord session: %v\n", err)
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!wgconfig" {
		// Implement Wireguard configuration provisioning here
		return
	}

}
