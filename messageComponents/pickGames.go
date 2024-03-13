package messageComponents

import (
	"MythiccBotHyper/globals"
	"MythiccBotHyper/model"
	"MythiccBotHyper/utils"
	"errors"
	"github.com/bwmarrin/discordgo"
	"log"
	"slices"
)

func pickGamesAdd(session *discordgo.Session, interaction *discordgo.InteractionCreate) {

	data, err := func() (*discordgo.InteractionResponseData, error) {
		// Get all games roles from model
		allGameRoles, err := model.GetAllSnowflakeIds(model.GameSnowflake{})
		if err != nil {
			log.Println(err)
			return nil, err
		}

		// Get game roles from user
		var userRoles []string
		if interaction.Member == nil {
			e := "Error: `interaction.Member` is nil"
			log.Println(e)
			return nil, err
		}
		userRoles = interaction.Member.Roles

		// Filter allGameRoles using userRoles
		allGameRoles = utils.Filter(allGameRoles, func(item string) bool {
			return !slices.Contains(userRoles, item)
		})

		// If there are any roles after filtering
		if len(allGameRoles) > 0 {
			minValue := 1

			// Get all roles from the guild
			guildRoles, err := session.GuildRoles(globals.GuildID)
			if err != nil {
				log.Println(err)
				return nil, err
			}

			// Map out role names and ids
			roleMap := make(map[string]string)
			for _, role := range guildRoles {
				if slices.Contains(allGameRoles, role.ID) {
					roleMap[role.Name] = role.ID
				}
			}

			// Build the options for the dropdown menu
			var options []discordgo.SelectMenuOption
			for key, value := range roleMap {
				options = append(options, discordgo.SelectMenuOption{
					Label: key,
					Value: value,
				})
			}

			data := &discordgo.InteractionResponseData{
				Content: "Please select the games you're interested in",
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.SelectMenu{
								CustomID:    "pick-games-add-execute",
								Placeholder: "No games selected",
								MinValues:   &minValue,
								MaxValues:   len(allGameRoles),
								Options:     options,
							},
						},
					},
				},
				Flags: discordgo.MessageFlagsEphemeral,
			}

			return data, nil
		}

		return nil, errors.New("you have already selected all available games")
	}()

	if err != nil {
		data = &discordgo.InteractionResponseData{
			Content: err.Error(),
			Flags:   discordgo.MessageFlagsEphemeral,
		}
	}

	err = session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: data,
	})
	if err != nil {
		log.Println(err)
	}

}

func pickGamesRemove(state *discordgo.Session, interaction *discordgo.InteractionCreate) {
	log.Println("pickGamesRemove executed")
}
