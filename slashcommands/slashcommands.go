package slashcommands

import (
	"MythiccBotHyper/globals"
	"github.com/bwmarrin/discordgo"
	"log"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "basic-command",
			Description: "Basic command",
		},
	}

	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"basic-command": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			globals.Bot.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Hey there! Congratulations, you just executed your first slash command",
				},
			})
		},
	}
)

func RegisterCommands() {
	log.Println("Registering commands...")
	currentCommands, err := globals.Bot.ApplicationCommands(globals.Bot.State.User.ID, globals.GuildID)
	if err != nil {
		log.Fatalf("Could not fetch registered commandsExample: %v", err)
	}

	for _, cmd := range Commands {
		skipCommand := false
		for _, currentCommand := range currentCommands {
			if cmd.Name == currentCommand.Name {
				skipCommand = true
				break
			}
		}

		if skipCommand {
			continue
		}

		_, err := globals.Bot.ApplicationCommandCreate(globals.Bot.State.User.ID, globals.GuildID, cmd)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", cmd.Name, err)
		}

		log.Println(cmd.Name, "was registered")

	}

	globals.Bot.AddHandler(func(bot *discordgo.Session, interactionCreate *discordgo.InteractionCreate) {
		if handler, ok := CommandHandlers[interactionCreate.ApplicationCommandData().Name]; ok {
			handler(bot, interactionCreate)
		}
	})
}

func UnregisterCommands() {
	if globals.RemoveCommands {
		log.Println("Removing commands...")
		registeredCommands, err := globals.Bot.ApplicationCommands(globals.Bot.State.User.ID, globals.GuildID)
		if err != nil {
			log.Fatalf("Could not fetch registered commandsExample: %v", err)
		}

		for _, v := range registeredCommands {
			log.Printf("Removing %v\n", v.Name)
			err := globals.Bot.ApplicationCommandDelete(globals.Bot.State.User.ID, globals.GuildID, v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}
}
