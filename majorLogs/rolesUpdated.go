package majorlogs

import (
	g "MythiccBotHyper/globals"
	"MythiccBotHyper/interactives"
	"MythiccBotHyper/utils"
	"fmt"
	"log"
	"slices"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func rolesUpdated(newRoles []string, removedRoles []string, user *discordgo.User, session *discordgo.Session) discordgo.MessageEmbed {
	fields := []*discordgo.MessageEmbedField{}

	timeStamp := time.Now().Format(time.RFC3339)
	url, _ := utils.GetAvatarUrl(user)
	userIdText := fmt.Sprintf("User ID: %v", user.ID)

	if len(newRoles) != 0 {
		guildApplyRoles := strings.Split(g.GuildApplyRoles, ",")
		setNeedsToApply := func() bool {
			for _, applyRole := range guildApplyRoles {
				if slices.Contains(newRoles, applyRole) {
					return true
				}
			}
			return false
		}()

		if setNeedsToApply {
			session.GuildMemberRoleAdd(g.GuildID, user.ID, g.NeedsToApplyRole)
		}

		if slices.Contains(newRoles, g.NeedsToApplyRole) {
			guildApplySendDirections(session, user)
		}

		for i, roleId := range newRoles {
			newRoles[i] = fmt.Sprintf("%v", interactives.FromRoleId(roleId))
		}
		val := strings.Join(newRoles, " ")
		fields = append(fields, &discordgo.MessageEmbedField{Name: "New Roles: ", Value: val, Inline: false})
	}

	if len(removedRoles) != 0 {
		for i, roleId := range removedRoles {
			removedRoles[i] = fmt.Sprintf("%v", interactives.FromRoleId(roleId))
		}
		val := strings.Join(removedRoles, " ")
		fields = append(fields, &discordgo.MessageEmbedField{Name: "Removed Roles: ", Value: val, Inline: false})
	}

	fields = append(fields, &discordgo.MessageEmbedField{Name: "Username", Value: interactives.FromUserId(user.ID), Inline: false})

	return discordgo.MessageEmbed{
		Title:       "Roles Updated",
		Color:       0xFEE75C,
		Description: "🔄 🔄 🔄",
		Fields:      fields,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    user.Username,
			IconURL: url,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: userIdText,
		},
		Timestamp: timeStamp,
	}
}

func guildApplySendDirections(session *discordgo.Session, user *discordgo.User) {
	channel, err := session.UserChannelCreate(user.ID)
	if err != nil {
		log.Println(err)
		return
	}

	// TODO: Make a NeedsToApplyHelpChannel global and use it
	msg := fmt.Sprintf("# Guild Application Required!\n%v",
		interactives.FromChannelId(g.NeedsToApplyChannel),
	)

	_, err = session.ChannelMessageSend(channel.ID, msg)
	if err != nil {
		log.Println(err)
		return
	}
}
