package utils

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
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

func DateDiff(date time.Time) (time.Time, error) {
	dateStr := date.String()
	dateStr = dateStr[0:10]
	joinDateArray := strings.Split(dateStr, "-")
	var joinDateIntArray []int

	for _, numStr := range joinDateArray {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return time.Time{}, err
		}
		joinDateIntArray = append(joinDateIntArray, num)
	}
	joinDate := time.Date(
		joinDateIntArray[0],
		time.Month(joinDateIntArray[1]),
		joinDateIntArray[2],
		0, 0, 0, 0, time.UTC,
	)
	currentDate := time.Now().UTC()
	currentDate = time.Date(
		currentDate.Year(),
		currentDate.Month(),
		20,
		0, 0, 0, 0, time.UTC,
	)

	days := int(currentDate.Sub(joinDate).Hours() / 24)
	years := int(days / 365)
	remaining_days := int(days % 365)
	months := int(remaining_days / 30)
	days = int(remaining_days % 30)

	return time.Date(
		years,
		time.Month(months),
		days,
		0, 0, 0, 0, time.UTC,
	), nil
}
