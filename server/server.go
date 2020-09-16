package server

import (
	"fmt"
	// "encoding/json"
	"log"
	"net/http"
	"runtime"

	"github.com/tesh254/thrusta/process"

	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"github.com/tesh254/thrusta/helpers"
	// "time"
)

// Channel handles channel type
type Channel struct {
	Channel string `json:"channel"`
}

// Message handles message type
type Message struct {
	ID      int    `json:"id"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func sendJoin(c *gosocketio.Client) {
	err := c.Emit("socket.io", helpers.ParseJSONFile())

	if err != nil {
		log.Fatal(err)
	}

	err = c.On("payload", func(h *gosocketio.Channel, args process.Processes) {
		var req process.Requests

		req.SocketProcess(args)
	})

	if err != nil {
		log.Println(err)
	}
}

// Server handles websocket server client
func Server() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wsURL string

	// if os.Getenv("THRUSTA_ENV") == "dev" {
	wsURL = gosocketio.GetUrl("localhost", 5000, false)
	// } else {
	// 	wsURL = gosocketio.GetUrl("172.17.0.6", 5000, false)
	// 	// wsURL = "wss://thrusta-api.app.bywachira.com"
	// }

	fmt.Println(wsURL)

	c, err := gosocketio.Dial(
		wsURL,
		transport.GetDefaultWebsocketTransport())

	if err != nil {
		log.Fatal(err)
	}

	err = c.On(gosocketio.OnConnection, func(h *gosocketio.Channel) {
		log.Println("🤝 Connected")
	})
	if err != nil {
		log.Println(err)
	}

	err = c.On(gosocketio.OnDisconnection, func(h *gosocketio.Channel) {
		log.Println("😔 Disconnected")
		log.Println("😬 Retrying in the next 10s")
		helpers.SetInterval(func() {
			log.Println("👋 Pinging the server...")
			err = c.On(gosocketio.OnConnection, func(h *gosocketio.Channel) {
				log.Println("🤝 Connected")
			})

			if err != nil {
				log.Println("😔 " + string(err.Error()))
			}
		}, 10)

	})
	if err != nil {
		log.Println(err)
	}

	helpers.SetInterval(func() { sendJoin(c) }, 5)
	serveMux := http.NewServeMux()
	serveMux.Handle("/node/", serveMux)

	log.Println("🏓 Starting server...")
	log.Panic(http.ListenAndServe(":4760", serveMux))
}
