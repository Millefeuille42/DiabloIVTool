package globals

import (
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
}
