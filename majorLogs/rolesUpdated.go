package majorlogs

import (
	"MythiccBotHyper/interactives"
	"MythiccBotHyper/utils"
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func RolesUpdated(newRoles []string, removedRoles []string, user discordgo.User) discordgo.MessageEmbed {
	fields := []*discordgo.MessageEmbedField{}

	timeStamp := time.Now().Format(time.RFC3339)
	url, _ := utils.GetAvatarUrl(&user)
	userIdText := fmt.Sprintf("User ID: %v", user.ID)

	if len(newRoles) != 0 {
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
