package redisCache

import "time"

type UniqueItemData struct {
	Name           string   `json:"name"`
	Class          string   `json:"class,omitempty"`
	Type           string   `json:"type"`
	SecondaryStats []string `json:"secondary_stats,omitempty"`
	TerciaryStats  []string `json:"terciary_stats,omitempty"`
	Effect         string   `json:"effect"`
	Flavor         string   `json:"flavor"`
	Filters        []string `json:"filters,omitempty"`
}

type UniqueItemsData struct {
	Items []UniqueItemData `json:"uniques"`
}

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
