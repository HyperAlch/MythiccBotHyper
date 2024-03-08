package slashcommands

import (
	"MythiccBotHyper/datatype"
	"MythiccBotHyper/globals"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

func prune(state *discordgo.Session, interaction *discordgo.InteractionCreate) {
	var contentMessage string
	contentMessage = func() string {
		// Pull user from interaction
		user, err := datatype.NewUserFromInteraction(interaction.Interaction)
		if err != nil {
			log.Println(err)
			return err.Error()
		}

		message := ""
		// Only admins are allowed to do this
		if !user.IsAdmin() {
			message = "You are not allowed to do this!"
		} else {
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
