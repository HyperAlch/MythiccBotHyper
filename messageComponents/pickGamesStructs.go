package messageComponents

import (
	"MythiccBotHyper/utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"slices"
	"strings"
	"time"
)

type GamesDropdown interface {
	Filter(item string, userRoles []string) bool
	GetContent() string
	GetCustomId() string
	GetDefaultMessage() string
}

type AddDropdown struct{}
type RemoveDropdown struct{}

func (_ AddDropdown) Filter(item string, userRoles []string) bool {
	return !slices.Contains(userRoles, item)
}

func (_ AddDropdown) GetContent() string {
	return "Please select the games you're interested in"
}

func (_ AddDropdown) GetCustomId() string {
	return "pick-games-add-execute"
}

func (_ AddDropdown) GetDefaultMessage() string {
	return "You have already selected all available games"
}

func (_ RemoveDropdown) Filter(item string, userRoles []string) bool {
	return slices.Contains(userRoles, item)
}

func (_ RemoveDropdown) GetContent() string {
	return "Please select the games you would like to remove"
}

func (_ RemoveDropdown) GetCustomId() string {
	return "pick-games-remove-execute"
}

func (_ RemoveDropdown) GetDefaultMessage() string {
	return "There are no games to remove"
}

type GamesDropdownExecute interface {
	ChangeUser(guildID, userID, roleID string, session *discordgo.Session) (err error)
	GetData(selectedRoles []string, user *discordgo.User) *discordgo.InteractionResponseData
}

type AddDropdownExecute struct{}
type RemoveDropdownExecute struct{}

func (_ AddDropdownExecute) ChangeUser(guildID, userID, roleID string, session *discordgo.Session) (err error) {
	err = session.GuildMemberRoleAdd(guildID, userID, roleID)
	if err != nil {
		return err
	}

	return nil
}

func (_ AddDropdownExecute) GetData(selectedRoles []string, user *discordgo.User) *discordgo.InteractionResponseData {
	return getEmbedData(":green_circle:   New Roles   :green_circle:", selectedRoles, user)
}

func (_ RemoveDropdownExecute) ChangeUser(guildID, userID, roleID string, session *discordgo.Session) (err error) {
	err = session.GuildMemberRoleRemove(guildID, userID, roleID)
	if err != nil {
		return err
	}

	return nil
}

func (_ RemoveDropdownExecute) GetData(selectedRoles []string, user *discordgo.User) *discordgo.InteractionResponseData {
	return getEmbedData(":red_circle:   Removed Roles   :red_circle:", selectedRoles, user)
}

func getEmbedData(displayTitle string, selectedRoles []string, user *discordgo.User) *discordgo.InteractionResponseData {
	for index, role := range selectedRoles {
		selectedRoles[index] = fmt.Sprintf("<@&%v>", role)
	}

	timeStamp := time.Now().Format(time.RFC3339)
	url, _ := utils.GetAvatarUrl(user)
	displayName := fmt.Sprintf("<@%v>", user.ID)
	userIdText := fmt.Sprintf("User ID: %v", user.ID)

	return &discordgo.InteractionResponseData{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title:       "Roles Updated",
				Color:       0xFEE75C,
				Description: "🔄 🔄 🔄",
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   displayTitle,
						Value:  strings.Join(selectedRoles, " "),
						Inline: true,
					},
					{
						Name:   "Display Name",
						Value:  displayName,
						Inline: false,
					},
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
		},
		Flags: discordgo.MessageFlagsEphemeral,
	}
}
