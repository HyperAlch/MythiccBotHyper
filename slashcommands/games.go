package slashcommands

import (
	"MythiccBotHyper/globals"
	"MythiccBotHyper/interactives"
	"MythiccBotHyper/model"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	gamesDetails = discordgo.ApplicationCommand{
		Name:        "games",
		Description: "Add, delete, or list games",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "list",
				Description: "List all games",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
			},
			{
				Name:        "add",
				Description: "Add a role to games",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionRole,
						Name:        "role",
						Description: "Game role",
						Required:    true,
					},
				},
			},
			{
				Name:        "remove",
				Description: "Remove a role from games",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionRole,
						Name:        "role",
						Description: "Game role",
						Required:    true,
					},
				},
			},
		},
	}
)

func games(state *discordgo.Session, interaction *discordgo.InteractionCreate) {
	contentMessage := ""

	options := interaction.ApplicationCommandData().Options
	selectedCommand := options[0].Name

	switch selectedCommand {
	case "list":
		contentMessage = gamesList()
	case "add":
		targetUser, err := getTargetRole(state, options)
		if err != nil {
			contentMessage = err.Error()
			break
		}
		contentMessage = gamesAdd(targetUser)
	case "remove":
		targetUser, err := getTargetRole(state, options)
		if err != nil {
			contentMessage = err.Error()
			break
		}
		contentMessage = gamesRemove(targetUser)
	default:
		contentMessage = "Invalid command"
	}

	err := globals.Bot.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: contentMessage,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	if err != nil {
		log.Println(err)
	}
}

func gamesList() string {
	allGames, err := model.GetAllSnowflakeIds(model.GameSnowflake{})
	if err != nil {
		log.Println(err)
		return err.Error()
	}

	for i, gamesId := range allGames {
		allGames[i] = interactives.FromRoleId(gamesId)
	}
	if len(allGames) < 1 {
		return "No games in database..."
	}

	return strings.Join(allGames, "\n")
}

func gamesRemove(user *discordgo.Role) string {
	id := user.ID
	err := model.RemoveSnowflakeById(id, model.GameSnowflake{})
	if err != nil {
		log.Println(err)
		return err.Error()
	}

	return fmt.Sprintf("%v was removed...", interactives.FromRoleId(id))
}

func gamesAdd(user *discordgo.Role) string {
	id := user.ID
	err := model.AddSnowflakeById(id, model.GameSnowflake{})
	if err != nil {
		log.Println(err)
		return err.Error()
	}

	return fmt.Sprintf("%v was added...", interactives.FromRoleId(id))
}

func getTargetRole(state *discordgo.Session, options []*discordgo.ApplicationCommandInteractionDataOption) (*discordgo.Role, error) {
	if options[0] != nil {
		if options[0].Options[0] != nil {
			if options[0].Options[0].RoleValue(state, globals.GuildID) != nil {
				return options[0].Options[0].RoleValue(state, globals.GuildID), nil
			}
		}
	}

	return nil, errors.New("target user missing from command or state")
}
