package slashcommands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	EsoFormDetails = discordgo.ApplicationCommand{
		Name:        "eso_form",
		Description: "Test the eso form",
	}
)

func EsoForm(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	err := session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,

		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "**Elder Scrolls Online** Guild Form",
					Color:       0x000000,
					Description: "Please fill out this short form to be added to the `Elder Scrolls Online` invite list!",
				},
			},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "I'm ready!",
							Style:    discordgo.SuccessButton,
							CustomID: "eso-form-execute",
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

func EsoForm_check_pc_or_console(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	err := session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,

		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: "Are you playing ESO on PC?",
					Color: 0x000000,
				},
			},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Yes, I'm on PC",
							Style:    discordgo.SuccessButton,
							CustomID: "eso-form-picked-pc",
						},
						discordgo.Button{
							Label:    "I'm on console",
							Style:    discordgo.SecondaryButton,
							CustomID: "eso-form-picked-console",
						},
					},
				},
			},
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		log.Println(err)
	}

}

func EsoForm_console_are_you_sure(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	err := session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,

		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: "Are you sure you are a console player?",
					Color: 0x000000,
				},
			},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Yes, I play on Playstation or Xbox",
							Style:    discordgo.DangerButton,
							CustomID: "eso-form-console-retard",
						},
						discordgo.Button{
							Label:    "I miss clicked, I'm on PC",
							Style:    discordgo.SecondaryButton,
							CustomID: "eso-form-picked-pc",
						},
					},
				},
			},
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		log.Println(err)
	}

}

func EsoForm_choose_content(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	minVal := 1
	menuOptions := []discordgo.SelectMenuOption{
		{
			Label: "PVP",
			Value: "PVP",
		},
		{
			Label: "PVE",
			Value: "PVE",
		},
		{
			Label: "Crafting",
			Value: "Crafting",
		},
		{
			Label: "Dailies",
			Value: "Dailies",
		},
	}

	err := session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,

		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: "What content are you interested in?",
					Color: 0x000000,
				},
			},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.SelectMenu{
							CustomID:    "eso-content-selected",
							Placeholder: "No content selected",
							MinValues:   &minVal,
							MaxValues:   len(menuOptions),
							Options:     menuOptions,
						},
					},
				},
			},
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		log.Println(err)
	}

}

func EsoForm_choose_party_roles(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	minVal := 1
	menuOptions := []discordgo.SelectMenuOption{
		{
			Label: "DPS",
			Value: "DPS",
		},
		{
			Label: "Healer",
			Value: "Healer",
		},
		{
			Label: "Tank",
			Value: "Tank",
		},
	}

	err := session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,

		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: "What do you plan to play?",
					Color: 0x000000,
				},
			},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.SelectMenu{
							CustomID:    "eso-party-role-selected",
							Placeholder: "No roles selected",
							MinValues:   &minVal,
							MaxValues:   len(menuOptions),
							Options:     menuOptions,
						},
					},
				},
			},
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		log.Println(err)
	}

}

func EsoForm_submit_name_buttons(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	err := session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,

		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "Character / Social Info",
					Color:       0x000000,
					Description: "Please submit both your characters name and @name",
				},
			},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Character Name",
							Style:    discordgo.PrimaryButton,
							CustomID: "eso-form-modal-character-name",
						},
						discordgo.Button{
							Label:    "Account Name (aka @name)",
							Style:    discordgo.SecondaryButton,
							CustomID: "eso-form-modal-account-name",
						},
						discordgo.Button{
							Label: "Help",
							Style: discordgo.LinkButton,
							URL:   "https://google.com",
						},
					},
				},
			},
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		log.Println(err)
	}

}

func EsoForm_submit_account_name(session *discordgo.Session, interaction *discordgo.InteractionCreate) {

	responseData := &discordgo.InteractionResponseData{
		CustomID: "eso_account_name_modal",
		Title:    "ESO Account Name",
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.TextInput{
						Label:     "Your EXACT account name (aka @name)",
						CustomID:  "eso-account-name-input",
						Style:     discordgo.TextInputShort,
						Required:  true,
						MaxLength: 32,
						MinLength: 1,
					},
				},
			},
		},
	}

	err := session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: responseData,
	})
	if err != nil {
		log.Println(err)
	}

}

func EsoForm_submit_character_name(session *discordgo.Session, interaction *discordgo.InteractionCreate) {

	responseData := &discordgo.InteractionResponseData{
		CustomID: "eso_character_name_modal",
		Title:    "ESO Character Name",
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.TextInput{
						Label:     "Your EXACT character name",
						CustomID:  "eso-character-name-input",
						Style:     discordgo.TextInputShort,
						Required:  true,
						MaxLength: 32,
						MinLength: 1,
					},
				},
			},
		},
	}

	err := session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: responseData,
	})
	if err != nil {
		log.Println(err)
	}

}
