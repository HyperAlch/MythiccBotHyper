package majorlogs

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func GuildMemberAdd(*discordgo.Session, *discordgo.GuildMemberAdd) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Bot Recovered:", r)
		}
	}()
	// TODO
	log.Println("Guild member joined...")
}
