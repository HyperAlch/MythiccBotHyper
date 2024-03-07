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

	if globals.DropTables {
		dropTables()
	}
	createTables()
}

func createTables() {
	createAdminsTables := `
	CREATE TABLE IF NOT EXISTS admins (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		snowflake TEXT NOT NULL UNIQUE
	)
	`

	_, err := DB.Exec(createAdminsTables)
	if err != nil {
		log.Fatal("Could not create admins table!")
	}

	checkMasterAdmin := "SELECT snowflake FROM admins WHERE snowflake = ?"
	row := DB.QueryRow(checkMasterAdmin, globals.MasterAdmin)

	var snowflake string
	_ = row.Scan(&snowflake)

	if snowflake == "" {
		insertMasterAdminQuery := "INSERT INTO admins(snowflake) VALUES (?)"
		_, err = DB.Exec(insertMasterAdminQuery, globals.MasterAdmin)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func dropTables() {
	dropAdminsTables := "DROP TABLE IF EXISTS admins"

	_, err := DB.Exec(dropAdminsTables)
	if err != nil {
		log.Println("Could not drop admins table!")
	}
}
