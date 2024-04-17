package messageComponents

import (
	"MythiccBotHyper/globals"
	"MythiccBotHyper/interactives"
	"errors"
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func pickGamesAddExecute(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	selectedRoles, data := pickGamesExecute(session, interaction, AddDropdownExecute{})
	guildApplyRoles := strings.Split(globals.GuildApplyRoles, ",")
	for i, role := range guildApplyRoles {
		guildApplyRoles[i] = interactives.FromRoleId(role)
		if slices.Contains(selectedRoles, guildApplyRoles[i]) {
			msg := fmt.Sprintf("# Guild Application Required!\n%v", interactives.FromChannelId(globals.NeedsToApplyChannel))
			data.Embeds = append(data.Embeds, &discordgo.MessageEmbed{
				Title:       "Guild Application Required!",
				Color:       0xED4245,
				Description: msg,
			})
		}
	}

	err := session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: data,
	})
	if err != nil {
		log.Println(err)
	}
}

func pickGamesRemoveExecute(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	_, data := pickGamesExecute(session, interaction, RemoveDropdownExecute{})

	err := session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: data,
	})
	if err != nil {
		log.Println(err)
	}
}

func pickGamesExecute(session *discordgo.Session, interaction *discordgo.InteractionCreate, dropdownExecute GamesDropdownExecute) (selectedRoles []string, data *discordgo.InteractionResponseData) {
	selectedRoles = interaction.MessageComponentData().Values
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

	return
}
