package db

import (
	"MythiccBotHyper/globals"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *sql.DB

func init() {
	db, err := sql.Open("sqlite3", "discord_bot.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	DB = db
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createDevTables()

	if globals.DropTables {
		dropTables()
	}
	createTables()
}

func createTables() {
	createAdminsTable()
	createGamesTable()
}

func dropTables() {
	dropAdminsTable()
	dropGamesTable()
}

func createDevTables() {
	createSnowflakeTable := `
	CREATE TABLE IF NOT EXISTS snowflake_table (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		snowflake TEXT NOT NULL UNIQUE
	)
	`
	_, _ = DB.Exec(createSnowflakeTable)
}
