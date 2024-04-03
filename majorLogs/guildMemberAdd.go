package majorlogs

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func GuildMemberAdd(*discordgo.Session, *discordgo.GuildMemberAdd) {
	// TODO
	log.Println("Guild member joined...")
}
