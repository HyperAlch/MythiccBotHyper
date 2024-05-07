package commands

import (
	applicationcommands "MythiccBotHyper/applicationCommands"
	"MythiccBotHyper/datatype"
	"MythiccBotHyper/globals"
	"MythiccBotHyper/messageComponents"
	"MythiccBotHyper/modalInteractions"
	"MythiccBotHyper/slashcommands"
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		//&pingDetails,
		&slashcommands.PruneDetails,
		&slashcommands.AdminsDetails,
		&slashcommands.GamesDetails,
		&slashcommands.PickGamesMenuDetails,
		&applicationcommands.GuildApplyDetails,
	}

	AdminCommandHandlers = datatype.InteractionMap{
		"ping":            slashcommands.Ping,
		"prune":           slashcommands.Prune,
		"admins":          slashcommands.Admins,
		"games":           slashcommands.Games,
		"pick_games_menu": slashcommands.PickGamesMenu,
	}

	CommandHandlers = datatype.InteractionMap{
		"Guild Apply":               applicationcommands.GuildApply,
		"guild_apply_modal":         modalInteractions.GuildApplyModal,
		"pick-games-add":            messageComponents.PickGamesAdd,
		"pick-games-remove":         messageComponents.PickGamesRemove,
		"pick-games-add-execute":    messageComponents.PickGamesAddExecute,
		"pick-games-remove-execute": messageComponents.PickGamesRemoveExecute,
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
