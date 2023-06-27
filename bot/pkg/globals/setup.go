package globals

import (
	"bot/pkg/utils"
	"os"
)

func SetGlobals() {
	DatabaseDriver = os.Getenv("DBIVTOOL_DB_DRIVER")
	if DatabaseDriver == "" {
		DatabaseDriver = "sqlite3"
	}

	DatabaseDSN = os.Getenv("DBIVTOOL_DB_DSN")
	if DatabaseDSN == "" {
		DatabaseDSN = "file:./db.sqlite3?_foreign_keys=ON"
	}

	BotToken = os.Getenv("DBIVTOOL_BOT_TOKEN")

	RedisHost = os.Getenv("DBIVTOOL_REDIS_HOST")
	if RedisHost == "" {
		RedisHost = "localhost"
	}

	RedisPort = os.Getenv("DBIVTOOL_REDIS_PORT")
	if RedisPort == "" {
		RedisPort = "6379"
	}

	RedisPassword = os.Getenv("DBIVTOOL_REDIS_PASSWORD")
	RedisDB = utils.SafeAtoi(os.Getenv("DBIVTOOL_REDIS_DB"))
}
