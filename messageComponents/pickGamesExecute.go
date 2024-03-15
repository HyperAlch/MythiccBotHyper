package messageComponents

import (
	"MythiccBotHyper/globals"
	"errors"
	"github.com/bwmarrin/discordgo"
	"log"
)

func pickGamesAddExecute(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	pickGamesExecute(session, interaction, AddDropdownExecute{})
}

func pickGamesRemoveExecute(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	pickGamesExecute(session, interaction, RemoveDropdownExecute{})
}

func pickGamesExecute(session *discordgo.Session, interaction *discordgo.InteractionCreate, dropdownExecute GamesDropdownExecute) {
	selectedRoles := interaction.MessageComponentData().Values
	var user *discordgo.User

	err := func() error {
		if interaction.Member == nil {
			return errors.New("interaction.Member is nil")
		}
		if interaction.Member.User == nil {
			return errors.New("interaction.Member.User is nil")
		}
		user = interaction.Member.User

		if len(selectedRoles) > 0 {
			for _, role := range selectedRoles {
				err := dropdownExecute.ChangeUser(globals.GuildID, user.ID, role, session)
				if err != nil {
					return err
				}
			}
		} else {
			return errors.New("role values missing from request")
		}

		return nil
	}()

	var data *discordgo.InteractionResponseData
	if err != nil {
		// Something failed, show the error
		data = &discordgo.InteractionResponseData{
			Content: err.Error(),
			Flags:   discordgo.MessageFlagsEphemeral,
		}
	} else {
		// Changes where successful, show the embed
		data = dropdownExecute.GetData(selectedRoles, user)
	}

	err = session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: data,
	})
	if err != nil {
		log.Println(err)
	}
}
