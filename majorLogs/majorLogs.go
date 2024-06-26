package majorlogs

import (
	"MythiccBotHyper/utils"
	"fmt"
	"slices"

	g "MythiccBotHyper/globals"

	"github.com/bwmarrin/discordgo"
)

// Triggered when a users Nickname or Roles change
func GuildMemberUpdate(session *discordgo.Session, guildMemberUpdate *discordgo.GuildMemberUpdate) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Bot Recovered:", r)
		}
	}()

	data := discordgo.MessageEmbed{}
	removedRoles := utils.Filter(guildMemberUpdate.BeforeUpdate.Roles, func(role string) bool {
		return !slices.Contains(guildMemberUpdate.Roles, role)
	})
	newRoles := utils.Filter(guildMemberUpdate.Roles, func(role string) bool {
		return !slices.Contains(guildMemberUpdate.BeforeUpdate.Roles, role)
	})

	if guildMemberUpdate.Nick != guildMemberUpdate.BeforeUpdate.Nick {
		data = nicknameUpdated(guildMemberUpdate.BeforeUpdate.Nick, guildMemberUpdate.Nick, guildMemberUpdate.User)
	} else if len(removedRoles) != 0 || len(newRoles) != 0 {
		data = rolesUpdated(newRoles, removedRoles, guildMemberUpdate.User, session)
	}

	isCached, index := g.CustomMembersState.Exists(guildMemberUpdate.User.ID)
	if isCached {
		go g.CustomMembersState.Update(index, *guildMemberUpdate.Member)
	} else {
		go g.CustomMembersState.Append(*guildMemberUpdate.Member)
	}

	_, err := session.ChannelMessageSendEmbed(
		g.MajorEventsChannel,
		&data,
	)
	if err != nil {
		return
	}
}
