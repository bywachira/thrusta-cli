package commandline

import (
	"context"
	"encoding/json"
	"log"
	"net/url"
	"time"

	"github.com/tesh254/thrusta-cli/config"
)

// RunWebSocketClient initialize the websocket client
func RunWebSocketClient() {
	_, cancel := context.WithCancel(context.Background())

	cliConfig := config.ReadConfig()

	protocol := config.ProtocolChecker()

	u := url.URL{Scheme: protocol, Host: cliConfig.URL, Path: cliConfig.Path}
	log.Printf("connecting to %s", u.String())

	var webSoc RecConn

	webSoc.Dial(u.String(), nil)

	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	for {
		select {
		// case <-ctx.Done():
		// 	go ws.Close()
		// 	log.Printf("Websocket closed: %s", ws.GetURL())
		// 	return
		default:
			if !webSoc.IsConnected() {
				log.Printf("Websocket disconnected %s", webSoc.url)
				continue
			}

			if err := webSoc.WriteMessage(1, []byte(sendCredentials())); err != nil {
				log.Printf("Error: WriteMessage %s", webSoc.url)
				return
			}

			_, message, err := webSoc.ReadMessage()

			if err != nil {
				log.Printf("Error: ReadMessage %s", webSoc.url)
				return
			}

			if message != nil {
				var req Requests

				var response Processes

				json.Unmarshal([]byte(message), &response)

				req.SocketProcess(response)
			}

			log.Printf("Success: %s", message)
		}
	}
}
