package majorlogs

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func GuildMemberUnbanned(*discordgo.Session, *discordgo.GuildBanRemove) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Bot Recovered:", r)
		}
	}()
	// TODO
	log.Println("Guild member unbanned...")
}
