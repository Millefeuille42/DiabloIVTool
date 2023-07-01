package redisCache

import (
	"context"
	"encoding/json"
)

func GetUpcomingHelltides() (UpcomingHelltidesData, error) {
	ctx := context.Background()
	data, err := Client.Get(ctx, "helltides").Bytes()
	if err != nil {
		return UpcomingHelltidesData{}, err
	}

	helltides := UpcomingHelltidesData{}
	err = json.Unmarshal(data, &helltides)
	if err != nil {
		return UpcomingHelltidesData{}, err
	}

	return helltides, nil
}

func GetUpcomingBosses() (UpcomingBossesData, error) {
	ctx := context.Background()
	data, err := Client.Get(ctx, "world_bosses").Bytes()
	if err != nil {
		return UpcomingBossesData{}, err
	}

	bosses := UpcomingBossesData{}
	err = json.Unmarshal(data, &bosses)
	if err != nil {
		return UpcomingBossesData{}, err
	}

	return bosses, nil
}

func GetHelltide() (HelltideData, error) {
	ctx := context.Background()
	data, err := Client.Get(ctx, "helltide").Bytes()
	if err != nil {
		return HelltideData{}, err
	}

	helltide := HelltideData{}
	err = json.Unmarshal(data, &helltide)
	if err != nil {
		return HelltideData{}, err
	}

	return helltide, nil
}

func GetWorldBoss() (WorldBossData, error) {
	ctx := context.Background()
	data, err := Client.Get(ctx, "world_boss").Bytes()
	if err != nil {
		return WorldBossData{}, err
	}

	worldBoss := WorldBossData{}
	err = json.Unmarshal(data, &worldBoss)
	if err != nil {
		return WorldBossData{}, err
	}

	return worldBoss, nil
}

func GetWorldBossZone() (string, error) {
	ctx := context.Background()
	data, err := Client.Get(ctx, "world_boss_zone").Bytes()
	if err != nil {
		return "", err
	}

	worldBossZone := ""
	err = json.Unmarshal(data, &worldBossZone)
	if err != nil {
		return "", err
	}

	return worldBossZone, nil
}

func GetUniqueItems() (UniqueItemsData, error) {
	ctx := context.Background()
	data, err := Client.Get(ctx, "unique_items").Bytes()
	if err != nil {
		return UniqueItemsData{}, err
	}

	uniqueItems := UniqueItemsData{}
	err = json.Unmarshal(data, &uniqueItems)
	if err != nil {
		return UniqueItemsData{}, err
	}

	return uniqueItems, nil
}
