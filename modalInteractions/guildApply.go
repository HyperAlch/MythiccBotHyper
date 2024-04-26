package modalInteractions

import (
	"MythiccBotHyper/globals"
	g "MythiccBotHyper/globals"
	"MythiccBotHyper/interactives"
	"fmt"
	"log"
	"slices"

	"github.com/bwmarrin/discordgo"
)

func GuildApplyModal(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	log.Println("Modal submit")

	nickname := interaction.
		ModalSubmitData().
		Components[0].(*discordgo.ActionsRow).
		Components[0].(*discordgo.TextInput).Value

	author := interaction.Interaction.Member
	hasRole := slices.Contains(author.Roles, g.NeedsToApplyRole)

	if hasRole {
		session.GuildMemberNickname(g.GuildID, author.User.ID, nickname)
		session.GuildMemberRoleRemove(g.GuildID, author.User.ID, g.NeedsToApplyRole)
		fields := []*discordgo.MessageEmbedField{
			{
				Name:   "Discord Username",
				Value:  author.User.Username,
				Inline: true,
			},
			{
				Name:   "Display Name",
				Value:  interactives.FromUserId(author.User.ID),
				Inline: true,
			},
			{
				Name:   "In-Game Name",
				Value:  nickname,
				Inline: false,
			},
		}
		embed := &discordgo.MessageEmbed{
			Title:       "Member Joined",
			Color:       0x57F287,
			Fields:      fields,
			Description: "Please DO NOT delete this after promoting or rejecting an applicant",
		}

		session.ChannelMessageSendEmbed(g.NeedsToApplyChannel, embed)

		err := globals.Bot.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Thanks %v!\n\nYour guild Application was sent!",
					interactives.FromUserId(author.User.ID)),
				Flags: discordgo.MessageFlagsEphemeral,
			},
		})

		if err != nil {
			log.Println("ERROR:", err)
		}
	}
}
