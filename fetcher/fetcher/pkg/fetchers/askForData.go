package fetchers

import (
	"fetcher/pkg/redisCache"
	"log"
)

func AskForData() {
	upcomingBosses, err := getUpcomingBossesRaw()
	if err != nil {
		log.Printf("Error getting upcoming bosses data: %v", err)
	} else {
		redisCache.Client.Set(redisCache.Context, "world_bosses", upcomingBosses, 0)
	}

	UpcomingHelltides, err := getUpcomingHelltidesRaw()
	if err != nil {
		log.Printf("Error getting upcoming helltides data: %v", err)
	} else {
		redisCache.Client.Set(redisCache.Context, "helltides", UpcomingHelltides, 0)
	}

	uniqueItems, err := getUniqueItemsRaw()
	if err != nil {
		log.Printf("Error getting unique items data: %v", err)
	} else {
		redisCache.Client.Set(redisCache.Context, "unique_items", uniqueItems, 0)
	}
}
