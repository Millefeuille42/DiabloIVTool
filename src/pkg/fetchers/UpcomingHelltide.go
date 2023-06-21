package fetchers

import (
	"diablo_iv_tool/pkg/globals"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const helltideEndpoint = "/helltides"
const helltideUrl = globals.ApiURL + helltideEndpoint

type UpcomingHelltideData struct {
	Time int `json:"time"`
}

func getUpcomingHelltideRaw() ([]byte, error) {
	resp, err := http.Get(helltideUrl)
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

func marshalUpcomingHelltide(data []byte) (UpcomingHelltideData, error) {
	upcoming := UpcomingHelltideData{}
	err := json.Unmarshal(data, &upcoming)
	if err != nil {
		log.Println(err)
		return UpcomingHelltideData{}, err
	}

	return upcoming, nil
}

func GetUpcomingHelltideData() (UpcomingHelltideData, error) {
	data, err := getUpcomingHelltideRaw()
	if err != nil {
		return UpcomingHelltideData{}, err
	}

	upcoming, err := marshalUpcomingHelltide(data)
	if err != nil {
		return UpcomingHelltideData{}, err
	}

	return upcoming, nil
}

func GetUpcomingHelltideFromStruct(upcoming UpcomingHelltideData) (int, int) {
	return upcoming.Time / 60, upcoming.Time % 60
}
