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

func GuildMemberRemove(session *discordgo.Session, guildMemberRemoveData *discordgo.GuildMemberRemove) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Bot Recovered:", r)
		}
	}()

	user := guildMemberRemoveData.User
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

	userRoles := []string{"Unknown"}

	if g.CustomMembersState.Length() > 0 {
		for i, member := range g.CustomMembersState.Members() {
			if member.User.ID == user.ID {
				if len(member.Roles) > 0 {
					userRoles = []string{}
					for _, role := range member.Roles {
						userRoles = append(userRoles, interactives.FromRoleId(role))
					}
				} else {
					userRoles[0] = "No roles"
				}
				go g.CustomMembersState.Delete(i)
				break
			}
		}
	}

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
			Title:       "Member Left",
			Color:       0xED4245,
			Fields:      fields,
			Description: interactives.FromUserId(user.ID),
			Image: &discordgo.MessageEmbedImage{
				URL: "https://i.ibb.co/1qyVmzG/left-discord.png",
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
