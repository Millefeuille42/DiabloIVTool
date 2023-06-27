package main

import (
	"fetcher/pkg/fetchers"
	"fetcher/pkg/globals"
	"fetcher/pkg/redisCache"
	"fetcher/pkg/wsFetcher"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

func askForData() {
	wsFetcher.AskForData(wsFetcher.Client.Send)
	fetchers.AskForData()
}

func main() {
	globals.SetGlobals()
	terminated := make(chan struct{}, 1)

	redisCache.Client = redisCache.New(&redis.Options{
		Addr:     globals.RedisHost + ":" + globals.RedisPort,
		Password: globals.RedisPassword,
		DB:       globals.RedisDB,
	})

	wsFetcher.Client = wsFetcher.New()
	resp, err := wsFetcher.Client.Connect()
	if err != nil {
		log.Printf("handshake failed with status %d", resp.StatusCode)
		log.Fatal("dial:", err)
	}
	wsFetcher.Client.Exited = terminated

	defer wsFetcher.Client.Close()
	go wsFetcher.Client.Listener()
	go wsFetcher.Client.Sender()

	updateTicker := time.NewTicker(time.Hour * 1)
	defer updateTicker.Stop()

	for {
		select {
		case <-updateTicker.C:
			log.Println("updating data")
		case <-terminated:
			log.Println("interrupt")
			return
		case <-wsFetcher.Client.Connected:
			log.Println("connected, asking for data")
			askForData()
		}
	}
}
