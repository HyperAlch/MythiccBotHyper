package model

import (
	"MythiccBotHyper/db"
	"database/sql"
	"log"
	"strings"
)

func GetAllSnowflakeIds(m SnowflakeModel) ([]string, error) {
	query := "SELECT snowflake FROM snowflake_table"
	if m != nil {
		query = strings.ReplaceAll(query, "snowflake_table", m.getTableName())
	}

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

func RemoveSnowflakeById(id string, m SnowflakeModel) error {
	query := "DELETE FROM snowflake_table WHERE snowflake = ?"
	if m != nil {
		query = strings.ReplaceAll(query, "snowflake_table", m.getTableName())
	}

	_, err := db.DB.Exec(query, id)
	return err
}

func AddSnowflakeById(id string, m SnowflakeModel) error {
	query := `
	INSERT INTO snowflake_table (snowflake)
	SELECT ?
	WHERE NOT EXISTS(SELECT 1 from snowflake_table WHERE snowflake = ?)
	`
	if m != nil {
		query = strings.ReplaceAll(query, "snowflake_table", m.getTableName())
	}

	_, err := db.DB.Exec(query, id, id)
	return err
}
