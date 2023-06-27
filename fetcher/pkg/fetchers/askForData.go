package fetchers

import (
	"context"
	"fetcher/pkg/redisCache"
	"log"
)

func AskForData() {
	ctx := context.Background()

	upcomingBosses, err := getUpcomingBossesRaw()
	if err != nil {
		log.Printf("Error getting upcoming bosses data: %v", err)
	} else {
		redisCache.Client.Set(ctx, "world_bosses", upcomingBosses, 0)
	}

	UpcomingHelltides, err := getUpcomingHelltidesRaw()
	if err != nil {
		log.Printf("Error getting upcoming helltides data: %v", err)
	} else {
		redisCache.Client.Set(ctx, "helltides", UpcomingHelltides, 0)
	}
}
