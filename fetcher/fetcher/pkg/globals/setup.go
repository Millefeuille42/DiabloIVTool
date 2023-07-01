package globals

import (
	"fetcher/pkg/utils"
	"os"
)

func SetGlobals() {
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
