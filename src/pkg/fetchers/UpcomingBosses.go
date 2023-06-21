package fetchers

import (
	"diablo_iv_tool/pkg/globals"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const worldBossesEndpoint = "/worldBossSpawns"
const worldBossesUrl = globals.ListApiURL + worldBossesEndpoint

type UpcomingBossesData struct {
	BossEvents []struct {
		Name string    `json:"boss"`
		Time time.Time `json:"time"`
	} `json:"bossEvents"`
}

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

func marshalUpcomingBosses(data []byte) (UpcomingBossesData, error) {
	upcoming := UpcomingBossesData{}
	err := json.Unmarshal(data, &upcoming)
	if err != nil {
		log.Println(err)
		return UpcomingBossesData{}, err
	}

	return upcoming, nil
}

func GetUpcomingBossesData() (UpcomingBossesData, error) {
	data, err := getUpcomingBossesRaw()
	if err != nil {
		return UpcomingBossesData{}, err
	}

	upcoming, err := marshalUpcomingBosses(data)
	if err != nil {
		return UpcomingBossesData{}, err
	}

	return upcoming, nil
}

func GetUpcomingBosses(loc *time.Location) (string, error) {
	upcoming, err := GetUpcomingBossesData()
	if err != nil {
		return "", err
	}

	bosses := ""

	for _, boss := range upcoming.BossEvents {
		date := boss.Time.In(loc).Format(time.RFC850)
		bosses = fmt.Sprintf("%s***%s***: %s\n", bosses, boss.Name, date)
	}

	return bosses, nil
}
