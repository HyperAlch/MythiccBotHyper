package applicationcommands

import "github.com/bwmarrin/discordgo"

var (
	guildApplyDetails = discordgo.ApplicationCommand{
		Name:        "Guild Apply",
		Description: "Apply for guild membership",
		Type:        discordgo.MessageApplicationCommand,
	}
)
