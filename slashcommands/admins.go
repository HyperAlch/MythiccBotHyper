package slashcommands

import (
	"MythiccBotHyper/datatype"
	"MythiccBotHyper/globals"
	"MythiccBotHyper/interactives"
	"MythiccBotHyper/model"
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

var (
	adminsDetails = discordgo.ApplicationCommand{
		Name:        "admins",
		Description: "Add, delete, or list admins",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "list",
				Description: "List all admins",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
			},
			{
				Name:        "add",
				Description: "Add a user to admins",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionUser,
						Name:        "user",
						Description: "User option",
						Required:    true,
					},
				},
			},
			{
				Name:        "remove",
				Description: "Remove a user from admins",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionUser,
						Name:        "user",
						Description: "User option",
						Required:    true,
					},
				},
			},
		},
	}
)

func admins(state *discordgo.Session, interaction *discordgo.InteractionCreate) {
	contentMessage := ""
	user, err := datatype.NewUserFromInteraction(interaction.Interaction)
	if err != nil {
		log.Println(err)
		contentMessage = err.Error()
	}

	if user.IsAdmin() {
		options := interaction.ApplicationCommandData().Options
		selectedCommand := options[0].Name

		switch selectedCommand {
		case "list":
			contentMessage = adminsList()
			break
		case "add":
			targetUser, err := getTargetUser(state, options)
			if err != nil {
				contentMessage = err.Error()
				break
			}
			contentMessage = adminsAdd(targetUser)
			break
		case "remove":
			targetUser, err := getTargetUser(state, options)
			if err != nil {
				contentMessage = err.Error()
				break
			}
			contentMessage = adminsRemove(targetUser)
			break
		default:
			contentMessage = "Invalid command"
		}
	} else {
		contentMessage = "You are not allowed to do this!"
	}

	err = globals.Bot.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
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

func adminsList() string {
	allAdmins, err := model.GetAllSnowflakeIds(model.AdminSnowflake{})
	if err != nil {
		log.Println(err)
		return err.Error()
	}

	for i, adminId := range allAdmins {
		allAdmins[i] = fmt.Sprintf("%v", interactives.FromUserId(adminId))
	}
	return strings.Join(allAdmins, "\n")
}

func adminsRemove(user *discordgo.User) string {
	id := user.ID
	u, err := datatype.NewUser(user)
	if err != nil {
		log.Println(err)
		return err.Error()
	}

	if u.IsMasterAdmin() {
		return "Nice try retard"
	}

	err = model.RemoveSnowflakeById(id, model.AdminSnowflake{})
	if err != nil {
		log.Println(err)
		return err.Error()
	}

	return fmt.Sprintf("%v was removed...", interactives.FromUserId(id))
}

func adminsAdd(user *discordgo.User) string {
	id := user.ID
	err := model.AddSnowflakeById(id, model.AdminSnowflake{})
	if err != nil {
		log.Println(err)
		return err.Error()
	}

	return fmt.Sprintf("%v was added...", interactives.FromUserId(id))
}

func getTargetUser(state *discordgo.Session, options []*discordgo.ApplicationCommandInteractionDataOption) (*discordgo.User, error) {
	if options[0] != nil {
		if options[0].Options[0] != nil {
			if options[0].Options[0].UserValue(state) != nil {
				return options[0].Options[0].UserValue(state), nil
			}
		}
	}

	return nil, errors.New("target user missing from command or state")
}
