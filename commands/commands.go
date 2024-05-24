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
		&slashcommands.EsoFormDetails,
	}

	AdminCommandHandlers = datatype.InteractionMap{
		"ping":            slashcommands.Ping,
		"prune":           slashcommands.Prune,
		"admins":          slashcommands.Admins,
		"games":           slashcommands.Games,
		"pick_games_menu": slashcommands.PickGamesMenu,
		"eso_form":        slashcommands.EsoForm,
		// TODO: Make a "Triggered!" applicationcommand
		// TODO: Make a "Release triggered!" applicationcommand
	}

	CommandHandlers = datatype.InteractionMap{
		"Guild Apply":                   applicationcommands.GuildApply,
		"guild_apply_modal":             modalInteractions.GuildApplyModal,
		"pick-games-add":                messageComponents.PickGamesAdd,
		"pick-games-remove":             messageComponents.PickGamesRemove,
		"pick-games-add-execute":        messageComponents.PickGamesAddExecute,
		"pick-games-remove-execute":     messageComponents.PickGamesRemoveExecute,
		"eso-form-execute":              slashcommands.EsoForm_check_pc_or_console,
		"eso-form-picked-console":       slashcommands.EsoForm_console_are_you_sure,
		"eso-form-picked-pc":            slashcommands.EsoForm_choose_party_roles,
		"eso-party-role-selected":       slashcommands.EsoForm_choose_content,
		"eso-content-selected":          slashcommands.EsoForm_submit_name_buttons,
		"eso-form-modal-character-name": slashcommands.EsoForm_submit_character_name,
		"eso-form-modal-account-name":   slashcommands.EsoForm_submit_account_name,
		"eso_account_name_modal": func(s *discordgo.Session, interaction *discordgo.InteractionCreate) {

			err := globals.Bot.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Account name submitted",
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})

			if err != nil {
				log.Println("ERROR:", err)
			}
		},
		"eso_character_name_modal": func(s *discordgo.Session, interaction *discordgo.InteractionCreate) {

			err := globals.Bot.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Character name submitted",
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})

			if err != nil {
				log.Println("ERROR:", err)
			}
		},
		"eso-form-console-retard": func(s *discordgo.Session, interaction *discordgo.InteractionCreate) {

			err := globals.Bot.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Get kicked console peasant!!!",
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})

			if err != nil {
				log.Println("ERROR:", err)
			}
		},
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
