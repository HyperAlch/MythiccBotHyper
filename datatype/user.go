package datatype

import (
	"MythiccBotHyper/db"
	"errors"
	"github.com/bwmarrin/discordgo"
)

type User struct {
	id string
}

func NewUser(user *discordgo.User) (User, error) {
	if user == nil {
		return User{}, errors.New("*discordgo.User is nil")
	}

	return User{
		id: user.ID,
	}, nil
}

func NewUserFromInteraction(interaction *discordgo.Interaction) (User, error) {
	if interaction != nil {
		if interaction.Member != nil {
			return NewUser(interaction.Member.User)
		}
	}

	return User{}, errors.New("could not find valid User in Interaction")
}

func (user User) Get() string {
	return user.id
}

func (user User) IsAdmin() bool {
	query := "SELECT snowflake FROM admins WHERE snowflake = ?"
	row := db.DB.QueryRow(query, user.Get())

	var result string
	_ = row.Scan(&result)

	if result == "" {
		return false
	}

	return true
}
