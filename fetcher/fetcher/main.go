package main

import (
	"fetcher/pkg/fetchers"
	"fetcher/pkg/globals"
	"fetcher/pkg/redisCache"
	"fetcher/pkg/wsFetcher"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"os/signal"
	"time"
)

func askForData(ws *wsFetcher.WsClient) {
	fetchers.AskForData()
	wsFetcher.AskForData(ws.Send)
}

func main() {
	globals.SetGlobals()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	redisCache.Context = redisCache.NewContext()

	redisCache.Client = redisCache.New(&redis.Options{
		Addr:         globals.RedisHost + ":" + globals.RedisPort,
		Password:     globals.RedisPassword,
		DB:           globals.RedisDB,
		ClientName:   "dbivtool-fetcher",
		MaxIdleConns: 5,
	})

	defer redisCache.Client.Close()

	ws := wsFetcher.New()
	resp, err := ws.Connect()
	if err != nil {
		log.Printf("main: handshake failed with status %d", resp.StatusCode)
		log.Fatal("main: dial:", err)
	}

	signal.Notify(ws.Interrupt, os.Interrupt)

	defer ws.Close()
	go ws.Listener()
	go ws.Sender()

	updateTicker := time.NewTicker(5 * time.Minute)
	defer updateTicker.Stop()

	for alive := true; alive; {
		select {
		case <-interrupt:
			log.Println("main: interrupt")
			alive = false
		case <-ws.Exited:
			log.Println("main: wsFetcher exited")
			alive = false
		case <-ws.Connected:
			log.Println("main: connected, asking for data")
			askForData(ws)
		case <-updateTicker.C:
			askForData(ws)
		}
	}
}
