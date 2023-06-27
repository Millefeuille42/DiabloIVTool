package wsFetcher

import (
	"github.com/gorilla/websocket"
)

type WsClient struct {
	done      chan struct{}
	conn      *websocket.Conn
	Send      chan string
	Connected chan struct{}
	Exited    chan struct{}
	Interrupt chan struct{}
}

func New() *WsClient {
	return &WsClient{
		done:      make(chan struct{}),
		Send:      make(chan string),
		Connected: make(chan struct{}),
		Exited:    make(chan struct{}),
		Interrupt: make(chan struct{}),
	}
}
