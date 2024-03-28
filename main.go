package main

import (
	"MythiccBotHyper/datatype"
	"MythiccBotHyper/db"
	g "MythiccBotHyper/globals"
	"MythiccBotHyper/messageComponents"
	"MythiccBotHyper/slashcommands"
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

	g.Bot.AddHandler(interactionCreate)
	g.Bot.AddHandler(voiceStateUpdate)

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
		return
	}

	slashcommands.RegisterCommands()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
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

func voiceStateUpdate(_ *discordgo.Session, voiceState *discordgo.VoiceStateUpdate) {
	beforeState := voiceState.BeforeUpdate
	currentState := voiceState

	if currentState == nil {
		log.Println("currentState is nil")
		return
	}

	if beforeState == nil {
		log.Println("User joined channel:", currentState.ChannelID)
	} else if beforeState.ChannelID != "" && currentState.ChannelID == "" {
		log.Println("User left channel:", beforeState.ChannelID)
	} else if beforeState.ChannelID != currentState.ChannelID {
		log.Printf("\nUser moved from %v to %v", beforeState.ChannelID, currentState.ChannelID)
	}

	// TODO: Log the voice channel changes via embeds
}
