package db

import (
	"MythiccBotHyper/globals"
	"log"
)

func createAdminsTable() {
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

	query := `
	INSERT INTO admins (snowflake)
	SELECT ?
	WHERE NOT EXISTS(SELECT 1 from admins WHERE snowflake = ?)
	`
	_, err = DB.Exec(query, globals.MasterAdmin, globals.MasterAdmin)
	if err != nil {
		log.Fatal(err)
	}
}

func dropAdminsTable() {
	dropAdminsTables := "DROP TABLE IF EXISTS admins"

	_, err := DB.Exec(dropAdminsTables)
	if err != nil {
		log.Println("Could not drop admins table!")
	}
}
