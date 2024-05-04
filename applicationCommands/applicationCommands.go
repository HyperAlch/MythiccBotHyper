package applicationcommands

import (
	"MythiccBotHyper/globals"
	"slices"

	"github.com/bwmarrin/discordgo"
)

var (
	GuildApplyDetails = discordgo.ApplicationCommand{
		Name: "Guild Apply",
		Type: discordgo.UserApplicationCommand,
	}
)

func GuildApply(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	hasApplyRole := slices.Contains(interaction.Member.Roles, globals.NeedsToApplyRole)

	responseType := discordgo.InteractionResponseModal
	responseData := &discordgo.InteractionResponseData{}

	if !hasApplyRole {
		responseType = discordgo.InteractionResponseChannelMessageWithSource
		responseData = &discordgo.InteractionResponseData{
			Content: "Application already sent OR you never selected a supported game while joining",
			Flags:   discordgo.MessageFlagsEphemeral,
		}
	} else {
		responseData = &discordgo.InteractionResponseData{
			CustomID: "guild_apply_modal",
			Title:    "Apply To Guild",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.TextInput{
							CustomID:    "in_game_name",
							Label:       "Your EXACT in-game name",
							Style:       discordgo.TextInputShort,
							Placeholder: "Sperg Warrior",
							Required:    true,
							MaxLength:   32,
							MinLength:   1,
						},
					},
				},
			},
		}
	}

	session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: responseType,
		Data: responseData,
	})
}
