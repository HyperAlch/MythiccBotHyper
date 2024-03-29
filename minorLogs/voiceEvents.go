package minorLogs

import (
	g "MythiccBotHyper/globals"
	"MythiccBotHyper/interactives"
	"MythiccBotHyper/utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

func VoiceStateUpdate(session *discordgo.Session, voiceState *discordgo.VoiceStateUpdate) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Bot Recovered:", r)
		}
	}()

	beforeState := voiceState.BeforeUpdate
	currentState := voiceState

	if currentState == nil {
		log.Println("currentState is nil")
		return
	}

	timeStamp := time.Now().Format(time.RFC3339)
	userIdText := fmt.Sprintf("User ID: %v", currentState.UserID)
	url, _ := utils.GetAvatarUrl(currentState.Member.User)
	embedTitle := ""
	embedColor := 0x000000
	embedDescription := ""

	fields := []*discordgo.MessageEmbedField{
		{
			Name:   "Display Name",
			Value:  interactives.FromUserId(currentState.UserID),
			Inline: false,
		},
	}

	if beforeState == nil {
		embedTitle = "Joined Voice Chat"
		embedColor = 0x57F287
		embedDescription = fmt.Sprintf("Channel: %v", interactives.FromChannelId(currentState.ChannelID))
	} else if beforeState.ChannelID != "" && currentState.ChannelID == "" {
		embedTitle = "Left Voice Chat"
		embedColor = 0xED4245
		embedDescription = fmt.Sprintf("Channel: %v", interactives.FromChannelId(beforeState.ChannelID))
	} else if beforeState.ChannelID != currentState.ChannelID {
		embedTitle = "Moved Voice Chat"
		embedColor = 0xFEE75C
		fields = []*discordgo.MessageEmbedField{
			{
				Name:   "Left",
				Value:  interactives.FromChannelId(beforeState.ChannelID),
				Inline: true,
			},
			{
				Name:   "Joined",
				Value:  interactives.FromChannelId(currentState.ChannelID),
				Inline: true,
			},
		}
	}

	data := &discordgo.MessageEmbed{
		Title:       embedTitle,
		Color:       embedColor,
		Description: embedDescription,
		Fields:      fields,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    currentState.Member.User.Username,
			IconURL: url,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: userIdText,
		},
		Timestamp: timeStamp,
	}

	_, err := session.ChannelMessageSendEmbed(
		g.MinorEventsChannel,
		data,
	)
	if err != nil {
		return
	}
}
