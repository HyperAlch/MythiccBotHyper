package majorlogs

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func GuildMemberRemove(*discordgo.Session, *discordgo.GuildMemberRemove) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Bot Recovered:", r)
		}
	}()
	// TODO
	log.Println("Guild member left...")
}
