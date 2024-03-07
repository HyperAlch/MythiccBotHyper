package globals

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	Token          string
	GuildID        string
	Bot            *discordgo.Session
	RemoveCommands bool
)

func init() {
	envVars, err := loadFromEnv("DISCORD_TOKEN", "GUILD_ID", "REMOVE_COMMANDS")
	if err != nil {
		log.Fatal(err)
	}

	Token = envVars["DISCORD_TOKEN"]
	GuildID = envVars["GUILD_ID"]
	RemoveCommands, err = strconv.ParseBool(envVars["REMOVE_COMMANDS"])
	if err != nil {
		log.Fatal(err)
	}

	// Create a new Discord session using the provided Bot token.
	Bot, err = discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
}

func loadFromEnv(keys ...string) (map[string]string, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("error loading .env file")
	}

	tokens := make(map[string]string)
	for _, key := range keys {
		token := os.Getenv(key)
		if token == "" {
			return nil, errors.New(fmt.Sprintf("%v not found", key))
		}
		tokens[key] = token
	}

	return tokens, nil
}
