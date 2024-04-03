package majorlogs

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func GuildMemberUnbanned(*discordgo.Session, *discordgo.GuildBanRemove) {
	// TODO
	log.Println("Guild member unbanned...")
}
