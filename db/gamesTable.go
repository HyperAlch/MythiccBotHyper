package db

import "log"

func createGamesTable() {
	createGamesTables := `
	CREATE TABLE IF NOT EXISTS games (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		snowflake TEXT NOT NULL UNIQUE
	)
	`

	_, err := DB.Exec(createGamesTables)
	if err != nil {
		log.Fatal("Could not create games table!")
	}
}

func dropGamesTable() {
	dropGamesTables := "DROP TABLE IF EXISTS games"

	_, err := DB.Exec(dropGamesTables)
	if err != nil {
		log.Println("Could not drop games table!")
	}
}
