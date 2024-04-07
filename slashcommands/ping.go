package slashcommands

import (
	"MythiccBotHyper/globals"
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	pingDetails = discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Check if bot is online",
	}
)

func ping(_ *discordgo.Session, interaction *discordgo.InteractionCreate) {
	err := globals.Bot.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Bot is online...",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	if err != nil {
		log.Println("ERROR:", err)
	}
}
