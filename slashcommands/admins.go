package slashcommands

import (
	"MythiccBotHyper/globals"
	"MythiccBotHyper/model"
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

	options := interaction.ApplicationCommandData().Options
	//jsonStr, _ := json.Marshal(options)
	//log.Println(string(jsonStr))
	selectedCommand := options[0].Name

	contentMessage := ""
	switch selectedCommand {
	case "list":
		contentMessage = adminsList()
		break
	case "add":
		contentMessage = "/admins add"
		break
	case "remove":
		contentMessage = "/admins remove"
		break
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

func adminsList() string {
	allAdmins, err := model.GetAllAdminIds()
	if err != nil {
		log.Println(err)
		return ""
	}

	for i, adminId := range allAdmins {
		allAdmins[i] = fmt.Sprintf("<@%v>", adminId)
	}
	return strings.Join(allAdmins, "\n")
}
