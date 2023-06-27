package wsFetcher

import (
	"fetcher/pkg/redisCache"
	"github.com/gorilla/websocket"
	"os"
)

type WsClient struct {
	conn          *websocket.Conn
	worldBossData redisCache.WorldBossData
	helltideData  redisCache.HelltideData
	isInitialized bool
	resetTimers   chan struct{}
	done          chan struct{}
	Send          chan string
	Connected     chan struct{}
	Exited        chan<- struct{}
	Interrupt     chan os.Signal
}

func New() *WsClient {
	return &WsClient{
		Send:          make(chan string),
		Interrupt:     make(chan os.Signal, 1),
		done:          make(chan struct{}),
		Connected:     make(chan struct{}),
		resetTimers:   make(chan struct{}),
		worldBossData: redisCache.WorldBossData{},
		helltideData:  redisCache.HelltideData{},
		isInitialized: false,
	}
}

var Client *WsClient
