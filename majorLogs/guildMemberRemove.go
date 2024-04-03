package majorlogs

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func GuildMemberRemove(*discordgo.Session, *discordgo.GuildMemberRemove) {
	// TODO
	log.Println("Guild member left...")
}
