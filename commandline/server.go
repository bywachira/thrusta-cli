package commandline

import (
	"encoding/json"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/tesh254/thrusta-cli/monitor"

	"github.com/tesh254/thrusta-cli/config"
	"github.com/tesh254/thrusta-cli/helpers"

	"github.com/gorilla/websocket"
)

func sendCredentials() string {
	creds, err := json.MarshalIndent(helpers.ParseJSONFile(), "", " ")

	if err != nil {
		log.Fatal("Please provide generate thrusta.json file")
	}

	return string(creds)
}

// RunServer connects and communicates with the API
func RunServer() {
	cliConfig := config.ReadConfig()

	protocol := config.ProtocolChecker()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: protocol, Host: cliConfig.URL, Path: cliConfig.Path}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	var req Requests

	helpers.Interval(func() {
		monitor.SendMonitorData()
	}, 300*time.Second)

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}

			var response Processes

			json.Unmarshal([]byte(message), &response)

			req.SocketProcess(response)
		}
	}()

	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(sendCredentials()))

			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
