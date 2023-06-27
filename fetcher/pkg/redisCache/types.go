package redisCache

import "time"

type UpcomingHelltidesData struct {
	HelltideEvents []time.Time `json:"helltideEvents"`
}

type UpcomingBossesData struct {
	BossEvents []struct {
		Name string    `json:"boss"`
		Time time.Time `json:"time"`
	} `json:"bossEvents"`
}

type HelltideData struct {
	EndTime   time.Time `json:"endTime"`
	StartTime time.Time `json:"startTime"`
}

type WorldBossData struct {
	Boss      string    `json:"boss"`
	StartTime time.Time `json:"startTime"`
}

type WorldBossZoneData string
