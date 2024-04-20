package slashcommands

import (
	"MythiccBotHyper/datatype"
	"MythiccBotHyper/globals"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	pruneAmount  = 1.0
	PruneDetails = discordgo.ApplicationCommand{
		Name:        "prune",
		Description: "Delete `x` amount of messages",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "amount",
				Description: "amount",
				MinValue:    &pruneAmount,
				MaxValue:    99,
				Required:    true,
			},
		},
	}
)

func Prune(state *discordgo.Session, interaction *discordgo.InteractionCreate) {
	contentMessage := func() string {
		message := ""
		// Only admins are allowed to do this
		// Pull `x`, the amount of messages to delete from interaction
		options := interaction.ApplicationCommandData().Options
		option := options[0]
		if option != nil {
			x := option.IntValue()
			message = fmt.Sprintf("Deleted %v messages...", x)

			// Get the IDs of messages to delete
			channelId := interaction.ChannelID
			messages, err := state.ChannelMessages(channelId, int(x), "", "", "")
			if err != nil {
				log.Println(err)
				message = "Could not query messages..."
			}

			var messageIds []string
			for _, message := range messages {
				messageIds = append(messageIds, message.ID)
			}

			// Delete the messages
			err = state.ChannelMessagesBulkDelete(channelId, messageIds)
			if err != nil {
				apiErr, isApiErr := datatype.NewApiError(err)
				if isApiErr {
					message = apiErr.Message
				}
				log.Println(err)
			}

		}
		return message
	}()

	err := globals.Bot.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: contentMessage,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	if err != nil {
		log.Println("ERROR:", err)
	}
}
