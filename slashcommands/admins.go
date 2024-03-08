package slashcommands

import (
	"MythiccBotHyper/globals"
	"github.com/bwmarrin/discordgo"
	"log"
)

var (
	adminsDetails = discordgo.ApplicationCommand{
		Name:        "admins",
		Description: "Add, delete, or list admins",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "list",
				Description: "List all admins",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
			},
			{
				Name:        "add",
				Description: "Add a user to admins",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionUser,
						Name:        "user",
						Description: "User option",
						Required:    true,
					},
				},
			},
			{
				Name:        "remove",
				Description: "Remove a user from admins",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionUser,
						Name:        "user",
						Description: "User option",
						Required:    true,
					},
				},
			},
		},
	}
)

func admins(state *discordgo.Session, interaction *discordgo.InteractionCreate) {
	err := globals.Bot.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "admin works",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	if err != nil {
		log.Println(err)
	}
}
