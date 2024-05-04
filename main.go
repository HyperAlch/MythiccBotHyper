package main

import (
	"MythiccBotHyper/commands"
	"MythiccBotHyper/datatype"
	"MythiccBotHyper/db"
	g "MythiccBotHyper/globals"
	majorlogs "MythiccBotHyper/majorLogs"
	"MythiccBotHyper/minorLogs"
	"database/sql"
	"fmt"
	"log"
	"maps"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func runProcess() (*int, error) {
	var pid int
	var system_process = &syscall.SysProcAttr{Noctty: true}
	var attr = os.ProcAttr{
		Dir: ".",
		Env: os.Environ(),
		Files: []*os.File{
			os.Stdin,
			nil,
			nil,
		},
		Sys: system_process,
	}
	process, err := os.StartProcess("/bin/sleep", []string{"sleep", "300"}, &attr)
	if err == nil {
		pid = process.Pid
		// It is not clear from docs, but Release actually detaches the process
		err = process.Release()
		if err != nil {
			return nil, err
		}

	} else {
		return nil, err
	}

	return &pid, nil
}

func isProcessRunning(pid int) bool {
	// Try to find the process by its PID
	process, err := os.FindProcess(pid)
	if err != nil {
		// If an error occurs, assume the process is not running
		return false
	}

	// Send signal 0 to check if the process exists
	err = process.Signal(os.Signal(syscall.Signal(0)))
	return err == nil
}

func getCommandFromPID(pid int) (string, error) {
	// Construct the path to the command line file
	cmdLinePath := fmt.Sprintf("/proc/%d/cmdline", pid)

	// Read the contents of the command line file
	cmdLineBytes, err := os.ReadFile(cmdLinePath)
	if err != nil {
		return "", err
	}

	// Convert the null-separated byte slice to a string
	// The command line arguments are separated by null bytes in the file
	cmdLine := string(cmdLineBytes)

	// Replace null bytes with spaces to make the command line more readable
	cmdLine = replaceNullBytesWithSpace(cmdLine)

	return cmdLine, nil
}

func replaceNullBytesWithSpace(s string) string {
	// Replace null bytes with spaces
	// This function helps in making the command line more readable
	output := strconv.QuoteToASCII(s)[1 : len(s)-1]
	output = strings.ReplaceAll(output, "\\x", " ")
	return output
}

func main() {
	// [Run a process]
	// pid, err := runProcess()
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// log.Println(*pid)

	// [Check if process is running]
	// alive := isProcessRunning(37144)
	// if alive {
	// 	log.Println("Process is alive")
	// }

	// [Get the name of the process]
	// cmdName, err := getCommandFromPID(45247)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// log.Println(cmdName)

	// startBot()
}

func startBot() {
	defer func(DB *sql.DB) {
		err := DB.Close()
		log.Println("")
		if err != nil {
			log.Println(err)
		} else {
			log.Println("SQL Database closed...")
		}
	}(db.DB)

	if g.Bot == nil {
		panic("Pointer to Bot is nil")
	}

	g.Bot.AddHandler(interactionCreate)

	// Minor Events
	g.Bot.AddHandler(minorLogs.VoiceStateUpdate)

	// Major Events
	g.Bot.AddHandler(majorlogs.GuildMemberUpdate)
	g.Bot.AddHandler(majorlogs.GuildMemberAdd)
	g.Bot.AddHandler(majorlogs.GuildMemberRemove)
	g.Bot.AddHandler(majorlogs.GuildMemberBanned)
	g.Bot.AddHandler(majorlogs.GuildMemberUnbanned)

	g.Bot.Identify.Intents = discordgo.IntentsGuilds |
		discordgo.IntentsGuildMessages |
		discordgo.IntentsMessageContent |
		discordgo.IntentsGuildVoiceStates |
		discordgo.IntentsGuildBans |
		discordgo.IntentsGuildPresences |
		discordgo.IntentsGuildMembers

	// Open a websocket connection to Discord and begin listening.
	err := g.Bot.Open()
	if err != nil {
		panic(err)
	}

	commands.RegisterCommands()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	commands.UnregisterCommands()

	// Cleanly close down the Discord session.
	defer func(Bot *discordgo.Session) {
		err := Bot.Close()
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Bot shutdown...")
		}
	}(g.Bot)

}

func interactionCreate(session *discordgo.Session, interactionCreate *discordgo.InteractionCreate) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Bot Recovered:", r)
		}
	}()

	executeInteraction := func(
		key string,
		user datatype.User,
	) {
		interactionMap := commands.CommandHandlers
		if user.IsAdmin() {
			interactionMap = commands.AdminCommandHandlers
			maps.Copy(interactionMap, commands.CommandHandlers)
		}

		handler, ok := interactionMap[key]
		if ok {
			handler(session, interactionCreate)
		}
	}

	interaction := interactionCreate.Interaction
	user, err := datatype.NewUserFromInteraction(interaction)
	if err != nil {
		log.Println("Could not get custom `datatype.User` from interaction")
		return
	}

	switch interactionCreate.Type {
	case discordgo.InteractionMessageComponent:
		executeInteraction(
			interactionCreate.MessageComponentData().CustomID,
			user,
		)
	case discordgo.InteractionApplicationCommand:
		executeInteraction(
			interactionCreate.ApplicationCommandData().Name,
			user,
		)
	case discordgo.InteractionModalSubmit:
		executeInteraction(
			interactionCreate.ModalSubmitData().CustomID,
			user,
		)
	default:
		log.Println("unknown interaction type:", interactionCreate.Type.String())
	}
}
