package main

import (
	"MythiccBotHyper/datatype"
	"MythiccBotHyper/db"
	g "MythiccBotHyper/globals"
	"MythiccBotHyper/messageComponents"
	"MythiccBotHyper/minorLogs"
	"MythiccBotHyper/slashcommands"
	"MythiccBotHyper/utils"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"slices"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	defer func(DB *sql.DB) {
		err := DB.Close()
		log.Println("")
		if err != nil {
			log.Println(err)
		} else {
			log.Println("SQL Database closed...")
		}
	}(db.DB)

	g.Bot.AddHandler(interactionCreate)
	g.Bot.AddHandler(minorLogs.VoiceStateUpdate)
	g.Bot.AddHandler(guildMemberUpdate)

	g.Bot.Identify.Intents = discordgo.IntentsGuilds |
		discordgo.IntentsGuildMessages |
		discordgo.IntentsMessageContent |
		discordgo.IntentsGuildVoiceStates |
		discordgo.IntentsGuildBans |
		discordgo.IntentsGuildPresences |
		discordgo.IntentsGuildMembers

	// Open a websocket connection to Discord and begin listening.
	err := g.Bot.Open()
	if err != nil {
		panic(err)
	}

	slashcommands.RegisterCommands()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	slashcommands.UnregisterCommands()

	// Cleanly close down the Discord session.
	defer func(Bot *discordgo.Session) {
		err := Bot.Close()
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Bot shutdown...")
		}
	}(g.Bot)

}

const messageComponent = 3 /* Button press, dropdown select, etc 	*/
const slashCommand = 2     /* Registered bot slash commands 		*/

func interactionCreate(session *discordgo.Session, interactionCreate *discordgo.InteractionCreate) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Bot Recovered:", r)
		}
	}()

	executeInteraction := func(key string, interactionMap datatype.InteractionMap) {
		handler, ok := interactionMap[key]
		if ok {
			handler(session, interactionCreate)
		}
	}

	if interactionCreate.Type == messageComponent {
		executeInteraction(
			interactionCreate.MessageComponentData().CustomID,
			messageComponents.MessageComponentHandlers,
		)
	} else if interactionCreate.Type == slashCommand {
		executeInteraction(
			interactionCreate.ApplicationCommandData().Name,
			slashcommands.CommandHandlers,
		)
	} else {
		log.Println("unknown interaction type:", interactionCreate.Type.String())
	}
}

// Triggered when a users Nickname or Roles change
func guildMemberUpdate(session *discordgo.Session, guildMemberUpdate *discordgo.GuildMemberUpdate) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Bot Recovered:", r)
		}
	}()

	log.Println("")
	removedRoles := utils.Filter(guildMemberUpdate.BeforeUpdate.Roles, func(role string) bool {
		return !slices.Contains(guildMemberUpdate.Roles, role)
	})
	addedRoles := utils.Filter(guildMemberUpdate.Roles, func(role string) bool {
		return !slices.Contains(guildMemberUpdate.BeforeUpdate.Roles, role)
	})

	if guildMemberUpdate.Nick != guildMemberUpdate.BeforeUpdate.Nick {
		log.Println("Nickname has changed...")
	} else if len(removedRoles) != 0 || len(addedRoles) != 0 {
		log.Println("Roles have changed...")
	}
}
