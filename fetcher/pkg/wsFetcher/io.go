package wsFetcher

import (
	"fetcher/pkg/redisCache"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

func askForHelltideData(sendChannel chan<- string) {
	sendChannel <- `{"t":"d","d":{"r":1,"a":"q","b":{"p":"/helltide","h":""}}}`
}

func askForWorldBossData(sendChannel chan<- string) {
	sendChannel <- `{"t":"d","d":{"r":1,"a":"q","b":{"p":"/world_boss","h":""}}}`
}

func askForWorldBossZoneData(sendChannel chan<- string) {
	sendChannel <- `{"t":"d","d":{"r":1,"a":"q","b":{"p":"/world_boss_zone","h":""}}}`
}

func AskForData(sendChannel chan<- string) {
	askForHelltideData(sendChannel)
	time.Sleep(1 * time.Second)
	askForWorldBossData(sendChannel)
	time.Sleep(1 * time.Second)
	askForWorldBossZoneData(sendChannel)
	time.Sleep(1 * time.Second)
}

func (client *WsClient) sendExited() {
	client.Exited <- struct{}{}
}

func (client *WsClient) Listener() {
	defer close(client.done)

	for {
		_, message, err := client.conn.ReadMessage()
		if err != nil {
			log.Println("wsFetcher: read:", err)
			return
		}
		log.Printf("wsFetcher: recv: %s", message)

		messageParsed, err := parseMessage(message)
		if err != nil {
			log.Println(err)
			continue
		}

		if messageParsed.Type == "c" {
			log.Println("wsFetcher: connected")
			client.Connected <- struct{}{}
			continue
		}

		err = client.parseMessageData(messageParsed, redisCache.Client)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}

func (client *WsClient) Sender() {
	ticker := time.NewTicker(time.Second * 50)
	defer ticker.Stop()

	for alive := true; alive; {
		select {
		case <-client.done:
			alive = false
		case m := <-client.Send:
			log.Printf("wsFetcher: send Message %s", m)
			err := client.conn.WriteMessage(websocket.TextMessage, []byte(m))
			if err != nil {
				log.Println("wsFetcher: write:", err)
				continue
			}
		case <-ticker.C:
			err := client.conn.WriteMessage(websocket.PingMessage, []byte("0"))
			if err != nil {
				log.Println("wsFetcher: write ping:", err)
				continue
			}
		case <-client.Interrupt:
			log.Println("wsFetcher: interrupt")
			err := client.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("wsFetcher: write close:", err)
				alive = false
			}
			select {
			case <-client.done:
			case <-time.After(time.Second):
			}
			alive = false
		}
	}
	client.sendExited()
}
