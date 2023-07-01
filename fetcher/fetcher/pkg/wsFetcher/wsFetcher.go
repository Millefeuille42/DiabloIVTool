package wsFetcher

import (
	"github.com/gorilla/websocket"
	"os"
)

type WsClient struct {
	done      chan struct{}
	conn      *websocket.Conn
	Send      chan string
	Connected chan struct{}
	Exited    chan struct{}
	Interrupt chan os.Signal
}

func New() *WsClient {
	return &WsClient{
		done:      make(chan struct{}),
		Send:      make(chan string),
		Connected: make(chan struct{}),
		Exited:    make(chan struct{}),
		Interrupt: make(chan os.Signal),
	}
}
