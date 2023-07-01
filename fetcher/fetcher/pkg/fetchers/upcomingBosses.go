package fetchers

import (
	"fetcher/pkg/globals"
	"io"
	"log"
	"net/http"
)

const worldBossesEndpoint = "/worldBossSpawns"
const worldBossesUrl = globals.ListApiURL + worldBossesEndpoint

func getUpcomingBossesRaw() ([]byte, error) {
	resp, err := http.Get(worldBossesUrl)
	if err != nil {
		return []byte(""), err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return []byte(""), err
	}

	return data, nil
}
