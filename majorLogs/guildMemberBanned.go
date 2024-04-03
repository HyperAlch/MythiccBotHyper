package majorlogs

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func GuildMemberBanned(*discordgo.Session, *discordgo.GuildBanAdd) {
	// TODO
	log.Println("Guild member banned...")
}
