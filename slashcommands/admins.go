package slashcommands

import (
	"MythiccBotHyper/datatype"
	"MythiccBotHyper/globals"
	"MythiccBotHyper/model"
	"encoding/json"
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
		jsonStr, _ := json.Marshal(options)
		log.Println(string(jsonStr))
		selectedCommand := options[0].Name

		switch selectedCommand {
		case "list":
			contentMessage = adminsList()
			break
		case "add":
			contentMessage = "/admins add"
			break
		case "remove":
			targetUser := options[0].Options[0].UserValue(state)
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
	allAdmins, err := model.GetAllAdminIds()
	if err != nil {
		log.Println(err)
		return err.Error()
	}

	for i, adminId := range allAdmins {
		allAdmins[i] = fmt.Sprintf("<@%v>", adminId)
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

	err = model.RemoveAdminById(id)
	if err != nil {
		log.Println(err)
		return err.Error()
	}

	return fmt.Sprintf("<@%v> was removed...", id)
}
