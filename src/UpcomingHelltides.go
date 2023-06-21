package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const helltidesEndpoint = "/helltides"
const helltidesUrl = listApiURL + helltidesEndpoint

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

func getUpcomingHelltidesData() (UpcomingHelltidesData, error) {
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

func getUpcomingHelltides() (string, error) {
	upcoming, err := getUpcomingHelltidesData()
	if err != nil {
		return "", err
	}

	Helltides := ""

	for _, Helltide := range upcoming.HelltideEvents {
		date := strings.Replace(Helltide.Local().Format(time.RFC850), " CEST", "", -1)
		Helltides = fmt.Sprintf("%s%s\n", Helltides, date)
	}

	return strings.Replace(Helltides, " CEST", "", -1), nil
}
