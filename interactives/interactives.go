package interactives

import "fmt"

func FromUserId(userId string) string {
	return fmt.Sprintf("<@%v>", userId)
}

func FromRoleId(roleId string) string {
	return fmt.Sprintf("<@&%v>", roleId)
}

func FromChannelId(channelId string) string {
	return fmt.Sprintf("<#%v>", channelId)
}
