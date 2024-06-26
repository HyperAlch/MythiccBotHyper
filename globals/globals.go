package globals

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	Token                    string
	GuildID                  string
	Bot                      *discordgo.Session
	MasterAdmin              string
	MinorEventsChannel       string
	MajorEventsChannel       string
	FollowerRole             string
	GuildApplyRoles          string
	NeedsToApplyRole         string
	NeedsToApplyChannel      string
	NeedsToApplyGuideChannel string
	CustomMembersState       *MembersState
	SyncSeconds              int
)

func init() {
	envVars, err := loadFromEnv("DISCORD_TOKEN",
		"GUILD_ID",
		"MASTER_ADMIN",
		"MINOR_EVENTS_CHANNEL",
		"MAJOR_EVENTS_CHANNEL",
		"FOLLOWER_ROLE",
		"GUILD_APPLY_ROLES",
		"NEEDS_TO_APPLY_ROLE",
		"NEEDS_TO_APPLY_CHANNEL",
		"NEEDS_TO_APPLY_GUIDE_CHANNEL",
		"SYNC_SECONDS",
	)
	if err != nil {
		log.Fatal(err)
	}

	CustomMembersState = &MembersState{
		members: []*discordgo.Member{},
	}

	Token = envVars["DISCORD_TOKEN"]
	GuildID = envVars["GUILD_ID"]
	MasterAdmin = envVars["MASTER_ADMIN"]
	MinorEventsChannel = envVars["MINOR_EVENTS_CHANNEL"]
	MajorEventsChannel = envVars["MAJOR_EVENTS_CHANNEL"]
	FollowerRole = envVars["FOLLOWER_ROLE"]
	GuildApplyRoles = envVars["GUILD_APPLY_ROLES"]
	NeedsToApplyRole = envVars["NEEDS_TO_APPLY_ROLE"]
	NeedsToApplyChannel = envVars["NEEDS_TO_APPLY_CHANNEL"]
	NeedsToApplyGuideChannel = envVars["NEEDS_TO_APPLY_GUIDE_CHANNEL"]
	SyncSeconds, err = strconv.Atoi(envVars["SYNC_SECONDS"])
	if err != nil {
		log.Panic(err)
	}

	// Create a new Discord session using the provided Bot token.
	Bot, err = discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatalln("error creating Discord session,", err)
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
			return nil, fmt.Errorf("%v not found", key)
		}
		tokens[key] = token
	}

	return tokens, nil
}
