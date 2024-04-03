package majorlogs

import (
	"MythiccBotHyper/interactives"
	"MythiccBotHyper/utils"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func nicknameUpdated(oldNickname string, newNickname string, user *discordgo.User) discordgo.MessageEmbed {
	if len(oldNickname) == 0 {
		oldNickname = "*Default Nickname*"
	}

	if len(newNickname) == 0 {
		newNickname = "*Default Nickname*"
	}

	fields := []*discordgo.MessageEmbedField{
		{
			Name:   "Old Nickname",
			Value:  oldNickname,
			Inline: true,
		},
		{
			Name:   "New Nickname",
			Value:  newNickname,
			Inline: true,
		},
		{
			Name:   "Username",
			Value:  interactives.FromUserId(user.ID),
			Inline: false,
		},
	}

	timeStamp := time.Now().Format(time.RFC3339)
	url, _ := utils.GetAvatarUrl(user)
	userIdText := fmt.Sprintf("User ID: %v", user.ID)
	return discordgo.MessageEmbed{
		Title:  "Nickname Changed",
		Color:  0xFEE75C,
		Fields: fields,
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
