package model

type SnowflakeModel interface {
	getTableName() string
}

type AdminSnowflake struct{}

func (_ AdminSnowflake) getTableName() string {
	return "admins"
}

type GameSnowflake struct{}

func (_ GameSnowflake) getTableName() string {
	return "games"
}
