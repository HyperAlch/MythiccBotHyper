package cliapp

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func CLIApp() {
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
							fmt.Println("starting detached bot...")
							// TODO
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
					fmt.Println("Checking if bot is running")
					// TODO
					return nil
				},
			},
			{
				Name:  "stop",
				Usage: "Stop the bot if running in the background",
				Action: func(ctx *cli.Context) error {
					fmt.Println("Stopping bot...")
					// TODO
					return nil
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
							// TODO
							return nil
						},
					},
					{
						Name:  "unregister",
						Usage: "Unregister all commands",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("Un-registering all commands...")
							// TODO
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
