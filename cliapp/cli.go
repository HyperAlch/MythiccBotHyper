package cliapp

import (
	botprocess "MythiccBotHyper/cliapp/botProcess"
	"MythiccBotHyper/commands"
	g "MythiccBotHyper/globals"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func CLIApp() {
	if _, err := os.Stat("./.pid"); errors.Is(err, os.ErrNotExist) {
		pidFileContent := []byte(fmt.Sprintf("%v", 0))
		err2 := os.WriteFile("./.pid", pidFileContent, 0644)
		if err2 != nil {
			panic(err2)
		}

	}

	if _, err := os.Stat("./.exeName"); errors.Is(err, os.ErrNotExist) {
		filename := []byte(filepath.Base(os.Args[0]))
		err2 := os.WriteFile("./.exeName", filename, 0644)
		if err2 != nil {
			panic(err2)
		}
	}

	app := &cli.App{
		Name:  "MythiccBot",
		Usage: "Manage your Mythicc Bot instance",
		Action: func(*cli.Context) error {
			fmt.Println("Argument required, use --help or -h for more details")
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:  "start",
				Usage: "Start the bot",
				Subcommands: []*cli.Command{
					{
						Name:  "detached",
						Usage: "Start as a background process",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("Starting detached bot...")
							botProcess, err := botprocess.SpawnProcess()
							if botProcess != nil {
								fmt.Println("Bot PID:", botProcess.PID())
							} else {
								return err
							}

							return nil
						},
					},
					{
						Name:  "attached",
						Usage: "Start bot directly in the current console",
						Action: func(cCtx *cli.Context) error {
							startBot()
							return nil
						},
					},
				},
			},
			{
				Name:  "check",
				Usage: "Check if the bot is running in the background",
				Action: func(ctx *cli.Context) error {
					botProcess, err := botprocess.FetchProcess()
					if err != nil {
						return err
					}

					if botProcess.IsRunning() {
						fmt.Println("Bot is running...")
					} else {
						fmt.Println("Bot is offline...")
					}

					return nil
				},
			},
			{
				Name:  "stop",
				Usage: "Stop the bot if running in the background",
				Action: func(ctx *cli.Context) error {
					botProcess, err := botprocess.FetchProcess()
					if err != nil {
						return err
					}

					if !botProcess.IsRunning() {
						fmt.Println("Bot is already offline!")
						return nil
					}

					err = botProcess.Stop()
					if err == nil {
						fmt.Println("Bot process was shutdown...")
					}
					return err
				},
			},
			{
				Name:  "database",
				Usage: "Perform a database action",
				Subcommands: []*cli.Command{
					{
						Name:  "reset",
						Usage: "Reset the current database and empty the tables",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("Resetting the database...")
							// TODO
							return nil
						},
					},
					{
						Name:  "backup",
						Usage: "Backup the current database",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("Backing up the current database...")
							// TODO
							return nil
						},
					},
				},
			},
			{
				Name:  "commands",
				Usage: "Manage the commands associated with the bot",
				Subcommands: []*cli.Command{
					{
						Name:  "register",
						Usage: "Register all commands",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("Registering all commands...")
							if g.Bot == nil {
								return errors.New("pointer to Bot is nil")
							}
							err := g.Bot.Open()
							if err != nil {
								return err
							}
							commands.RegisterCommands()
							g.Bot.Close()
							return nil
						},
					},
					{
						Name:  "unregister",
						Usage: "Unregister all commands",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("Un-registering all commands...")
							if g.Bot == nil {
								return errors.New("pointer to Bot is nil")
							}
							err := g.Bot.Open()
							if err != nil {
								return err
							}
							commands.UnregisterCommands()
							g.Bot.Close()
							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
