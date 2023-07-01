package wsFetcher

import (
	"encoding/json"
	"fetcher/pkg/redisCache"
	"github.com/redis/go-redis/v9"
	"time"
)

func parseHelltideData(data []byte) (redisCache.HelltideData, error) {
	helltide := redisCache.HelltideData{}
	err := json.Unmarshal(data, &helltide)
	if err != nil {
		return redisCache.HelltideData{}, err
	}

	return helltide, nil
}

func parseWorldBossData(data []byte) (redisCache.WorldBossData, error) {
	worldBoss := redisCache.WorldBossData{}
	err := json.Unmarshal(data, &worldBoss)
	if err != nil {
		return redisCache.WorldBossData{}, err
	}

	return worldBoss, nil
}

func parseWorldBossZoneData(data []byte) (redisCache.WorldBossZoneData, error) {
	var worldBossZone redisCache.WorldBossZoneData = ""
	err := json.Unmarshal(data, &worldBossZone)
	if err != nil {
		return "", err
	}

	return worldBossZone, nil
}

func (client *WsClient) parseMessageData(message socketData, redisClient *redis.Client) error {
	dataJson, err := json.Marshal(message.Data.Body.Data)
	if err != nil {
		return err
	}

	switch message.Data.Body.Title {
	case "helltide":
		_, err = parseHelltideData(dataJson)
		if err != nil {
			return err
		}
	case "world_boss":
		_, err = parseWorldBossData(dataJson)
		if err != nil {
			return err
		}
	case "world_boss_zone":
		_, err = parseWorldBossZoneData(dataJson)
		if err != nil {
			return err
		}
	default:
		return nil
	}

	err = redisClient.Set(redisCache.Context, message.Data.Body.Title, dataJson, time.Minute*30).Err()
	if err != nil {
		return err
	}

	return nil
}

func parseMessage(message []byte) (socketData, error) {
	messageParsed := socketData{}
	err := json.Unmarshal(message, &messageParsed)
	if err != nil {
		return socketData{}, err
	}

	return messageParsed, nil
}
