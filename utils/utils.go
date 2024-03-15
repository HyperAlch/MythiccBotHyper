package utils

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"net/http"
)

func Filter[T any](items []T, fn func(item T) bool) []T {
	var filteredItems []T
	for _, value := range items {
		if fn(value) {
			filteredItems = append(filteredItems, value)
		}
	}
	return filteredItems
}

func GetAvatarUrl(user *discordgo.User) (string, error) {
	avatarUrl := user.Avatar
	userId := user.ID

	url := fmt.Sprintf("https://cdn.discordapp.com/avatars/%v/%v.png", userId, avatarUrl)
	_, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return url, nil
}
