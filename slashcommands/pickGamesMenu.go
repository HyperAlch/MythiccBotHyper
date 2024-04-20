package slashcommands

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	PickGamesMenuDetails = discordgo.ApplicationCommand{
		Name:        "pick_games_menu",
		Description: "Setup the `Pick Your Games` menu",
	}
)

func PickGamesMenu(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	var instructions = make([]string, 3)

	instructions[0] = ":green_circle: **Add** - Press to get a dropdown of all available game roles that you don't already have, select the ones you want.\n\n"
	instructions[1] = ":red_circle: **Remove** - Press to get a dropdown of all game roles currently assigned to you, select the ones you want to remove.\n\n"
	instructions[2] = "**[IMPORTANT]\n\nDropdowns are meant for ONE TIME USE ONLY. Please press \"Dismiss message\" when you are done. \n\nThese temporary messages DO NOT automatically update to reflect role changes, you must press the buttons AGAIN to make any changes in the future!**"

	err := session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "# ~ Pick Your Games ~",
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "Instructions",
					Color:       0x000000,
					Description: strings.Join(instructions, ""),
				},
			},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Add",
							Style:    discordgo.SuccessButton,
							CustomID: "pick-games-add",
						},
						discordgo.Button{
							Label:    "Remove",
							Style:    discordgo.DangerButton,
							CustomID: "pick-games-remove",
						},
					},
				},
			},
		},
	})
	if err != nil {
		log.Println(err)
	}
}
