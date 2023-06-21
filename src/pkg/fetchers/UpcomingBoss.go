package fetchers

import (
	"diablo_iv_tool/pkg/globals"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const worldBossEndpoint = "/world-bosses"
const worldBossUrl = globals.ApiURL + worldBossEndpoint

type UpcomingBossData struct {
	Name string `json:"name"`
	Time int    `json:"time"`
}

func getUpcomingBossRaw() ([]byte, error) {
	resp, err := http.Get(worldBossUrl)
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

func marshalUpcomingBoss(data []byte) (UpcomingBossData, error) {
	upcoming := UpcomingBossData{}
	err := json.Unmarshal(data, &upcoming)
	if err != nil {
		log.Println(err)
		return UpcomingBossData{}, err
	}

	return upcoming, nil
}

func GetUpcomingBossData() (UpcomingBossData, error) {
	data, err := getUpcomingBossRaw()
	if err != nil {
		return UpcomingBossData{}, err
	}

	upcoming, err := marshalUpcomingBoss(data)
	if err != nil {
		return UpcomingBossData{}, err
	}

	return upcoming, nil
}

func GetUpcomingBossFromStruct(upcoming UpcomingBossData) (string, int, int) {
	return upcoming.Name, upcoming.Time / 60, upcoming.Time % 60
}
