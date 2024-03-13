package datatype

import "github.com/bwmarrin/discordgo"

type InteractionMap map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
