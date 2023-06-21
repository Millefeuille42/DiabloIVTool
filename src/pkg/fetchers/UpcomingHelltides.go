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

const helltidesEndpoint = "/helltides"
const helltidesUrl = globals.ListApiURL + helltidesEndpoint

type UpcomingHelltidesData struct {
	HelltideEvents []time.Time `json:"helltideEvents"`
}

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

func marshalUpcomingHelltides(data []byte) (UpcomingHelltidesData, error) {
	upcoming := UpcomingHelltidesData{}
	err := json.Unmarshal(data, &upcoming)
	if err != nil {
		log.Println(err)
		return UpcomingHelltidesData{}, err
	}

	return upcoming, nil
}

func GetUpcomingHelltidesData() (UpcomingHelltidesData, error) {
	data, err := getUpcomingHelltidesRaw()
	if err != nil {
		return UpcomingHelltidesData{}, err
	}

	upcoming, err := marshalUpcomingHelltides(data)
	if err != nil {
		return UpcomingHelltidesData{}, err
	}

	return upcoming, nil
}

func GetUpcomingHelltides(loc *time.Location) (string, error) {
	upcoming, err := GetUpcomingHelltidesData()
	if err != nil {
		return "", err
	}

	Helltides := ""

	for _, Helltide := range upcoming.HelltideEvents {
		date := Helltide.In(loc).Format(time.RFC850)
		Helltides = fmt.Sprintf("%s%s\n", Helltides, date)
	}

	return Helltides, nil
}
