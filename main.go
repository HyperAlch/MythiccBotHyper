package main

import (
	"MythiccBotHyper/commands"
	"MythiccBotHyper/datatype"
	"MythiccBotHyper/db"
	g "MythiccBotHyper/globals"
	majorlogs "MythiccBotHyper/majorLogs"
	"MythiccBotHyper/messageComponents"
	"MythiccBotHyper/minorLogs"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
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

	if g.Bot == nil {
		panic("Pointer to Bot is nil")
	}

	g.Bot.AddHandler(interactionCreate)

	// Minor Events
	g.Bot.AddHandler(minorLogs.VoiceStateUpdate)

	// Major Events
	g.Bot.AddHandler(majorlogs.GuildMemberUpdate)
	g.Bot.AddHandler(majorlogs.GuildMemberAdd)
	g.Bot.AddHandler(majorlogs.GuildMemberRemove)
	g.Bot.AddHandler(majorlogs.GuildMemberBanned)
	g.Bot.AddHandler(majorlogs.GuildMemberUnbanned)

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

	commands.RegisterCommands()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	commands.UnregisterCommands()

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

const messageComponentValue = 3 /* Button press, dropdown select, etc 	*/
const slashCommandValue = 2     /* Registered bot slash commands 		*/

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

	if interactionCreate.Type == messageComponentValue {
		executeInteraction(
			interactionCreate.MessageComponentData().CustomID,
			messageComponents.MessageComponentHandlers,
		)
	} else if interactionCreate.Type == slashCommandValue {
		interaction := interactionCreate.Interaction
		user, err := datatype.NewUserFromInteraction(interaction)
		if err != nil {
			log.Println("Could not get custom `datatype.User` from interaction")
			return
		}

		if !user.IsAdmin() {
			log.Printf("Non-admin user %v tried to execute a slash command\n", user.Get())
			return
		}

		executeInteraction(
			interactionCreate.ApplicationCommandData().Name,
			commands.CommandHandlers,
		)
	} else {
		log.Println("unknown interaction type:", interactionCreate.Type.String())
	}
}
