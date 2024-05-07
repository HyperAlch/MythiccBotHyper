package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
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
	createTables()
}

func createTables() {
	createAdminsTable()
	createGamesTable()
}

func DropTables() {
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
