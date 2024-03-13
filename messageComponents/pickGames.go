package messageComponents

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func pickGamesAdd(state *discordgo.Session, interaction *discordgo.InteractionCreate) {
	log.Println("pickGamesAdd executed")
}

func pickGamesRemove(state *discordgo.Session, interaction *discordgo.InteractionCreate) {
	log.Println("pickGamesRemove executed")
}
