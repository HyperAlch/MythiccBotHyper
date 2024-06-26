package majorlogs

import (
	g "MythiccBotHyper/globals"
	"MythiccBotHyper/interactives"
	"MythiccBotHyper/utils"
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func GuildMemberBanned(session *discordgo.Session, guildMemberBannedData *discordgo.GuildBanAdd) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Bot Recovered:", r)
		}
	}()

	user := guildMemberBannedData.User
	joinDate, err := discordgo.SnowflakeTimestamp(user.ID)
	if err != nil {
		log.Println(err)
		return
	}

	dateDiff, err := utils.DateDiff(joinDate)
	if err != nil {
		log.Println(err)
		return
	}

	years := dateDiff.Year()
	months := int(dateDiff.Month())
	days := int(dateDiff.Day())

	formattedDateDiff := fmt.Sprintf(
		"**%v** ***years*** | **%v** ***months*** | **%v** ***days***",
		years, months, days,
	)

	fields := []*discordgo.MessageEmbedField{
		{
			Name:   "Account Age",
			Value:  formattedDateDiff,
			Inline: true,
		},
	}

	url, _ := utils.GetAvatarUrl(user)
	userIdText := fmt.Sprintf("User ID: %v", user.ID)
	timeStamp := time.Now().Format(time.RFC3339)

	_, err = session.ChannelMessageSendEmbed(
		g.MajorEventsChannel,
		&discordgo.MessageEmbed{
			Title:       "Member Banned",
			Color:       0xED4245,
			Fields:      fields,
			Description: interactives.FromUserId(user.ID),
			Image: &discordgo.MessageEmbedImage{
				URL: "https://i.ibb.co/P4m8YSL/banned.png",
			},
			Author: &discordgo.MessageEmbedAuthor{
				Name:    user.Username,
				IconURL: url,
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: userIdText,
			},
			Timestamp: timeStamp,
		},
	)
	if err != nil {
		return
	}
}
