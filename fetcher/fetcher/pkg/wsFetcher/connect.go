package wsFetcher

import (
	"fetcher/pkg/globals"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
)

func (client *WsClient) Connect() (*http.Response, error) {
	u := url.URL{
		Scheme:   "wss",
		Host:     globals.WsApiHost,
		Path:     globals.WsApiPath,
		RawQuery: globals.WsApiQuery,
	}
	log.Printf("wsFetcher: connecting to %s", u.String())

	conn, resp, err := websocket.DefaultDialer.Dial(u.String(), nil)
	client.conn = conn

	return resp, err
}

func (client *WsClient) Close() {
	client.conn.Close()
}
