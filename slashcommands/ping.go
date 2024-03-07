package slashcommands

import (
	"MythiccBotHyper/globals"
	"github.com/bwmarrin/discordgo"
	"log"
)

func ping(_ *discordgo.Session, i *discordgo.InteractionCreate) {
	err := globals.Bot.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Bot is online...",
		},
	})

	if err != nil {
		log.Println("ERROR:", err)
	}
}
