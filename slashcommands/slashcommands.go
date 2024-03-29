package slashcommands

import (
	"MythiccBotHyper/datatype"
	"MythiccBotHyper/globals"
	"github.com/bwmarrin/discordgo"
	"log"
)

var (
	pruneAmount = 1.0
	Commands    = []*discordgo.ApplicationCommand{
		//&pingDetails,
		&pruneDetails,
		&adminsDetails,
		&gamesDetails,
		&pickGamesMenuDetails,
	}

	CommandHandlers = datatype.InteractionMap{
		"ping":            ping,
		"prune":           prune,
		"admins":          admins,
		"games":           games,
		"pick_games_menu": pickGamesMenu,
	}
)

func RegisterCommands() {
	log.Println("Registering commands...")

	// Get all currently registered commands
	currentCommands, err := globals.Bot.ApplicationCommands(globals.Bot.State.User.ID, globals.GuildID)
	if err != nil {
		log.Fatalf("Could not fetch registered commandsExample: %v", err)
	}

	// Loop through the commands we want to register
	for _, cmd := range Commands {

		// Skip commands that have already been registered
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

		// Register any new commands
		_, err := globals.Bot.ApplicationCommandCreate(globals.Bot.State.User.ID, globals.GuildID, cmd)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", cmd.Name, err)
		}

		log.Println(cmd.Name, "was registered")
	}

}
func UnregisterCommands() {
	if globals.RemoveCommands {
		log.Println("Removing commands...")

		// Get all registered commands
		registeredCommands, err := globals.Bot.ApplicationCommands(globals.Bot.State.User.ID, globals.GuildID)
		if err != nil {
			log.Fatalf("Could not fetch registered commandsExample: %v", err)
		}

		// Remove all the commands
		for _, command := range registeredCommands {
			log.Printf("Removing %v\n", command.Name)
			err := globals.Bot.ApplicationCommandDelete(globals.Bot.State.User.ID, globals.GuildID, command.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", command.Name, err)
			}
		}
	}
}
