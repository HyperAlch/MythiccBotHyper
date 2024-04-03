package majorlogs

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func GuildMemberBanned(*discordgo.Session, *discordgo.GuildBanAdd) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Bot Recovered:", r)
		}
	}()
	// TODO
	log.Println("Guild member banned...")
}
