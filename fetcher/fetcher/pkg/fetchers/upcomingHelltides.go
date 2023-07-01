package fetchers

import (
	"fetcher/pkg/globals"
	"io"
	"log"
	"net/http"
)

const helltidesEndpoint = "/helltides"
const helltidesUrl = globals.ListApiURL + helltidesEndpoint

func getUpcomingHelltidesRaw() ([]byte, error) {
	resp, err := http.Get(helltidesUrl)
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
