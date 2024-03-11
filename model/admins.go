package model

import (
	"MythiccBotHyper/db"
	"database/sql"
	"log"
)

func GetAllAdminIds() ([]string, error) {
	query := "SELECT snowflake FROM admins"

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}(rows)

	var users []string

	for rows.Next() {
		var snowflake string
		err := rows.Scan(&snowflake)
		if err != nil {
			return nil, err
		}

		if snowflake == "" {
			continue
		}

		users = append(users, snowflake)
	}

	return users, nil
}

func RemoveAdminById(id string) error {
	query := "DELETE FROM admins WHERE snowflake = ?"
	_, err := db.DB.Exec(query, id)
	return err
}

func AddAdminById(id string) error {
	query := `
	INSERT INTO admins (snowflake)
	SELECT ?
	WHERE NOT EXISTS(SELECT 1 from admins WHERE snowflake = ?)
	`
	_, err := db.DB.Exec(query, id, id)
	return err
}
