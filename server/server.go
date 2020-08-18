package server

import (
	"log"
	"net/http"
	"runtime"

	"github.com/graarh/golang-socketio"
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
	err := c.Emit("socket.io", Channel{"main"})

	if err != nil {
		log.Fatal(err)
	}

	err = c.On("payload", func(h *gosocketio.Channel, args Channel) {
		log.Println(args)
	})

	if err != nil {
		log.Fatal(err)
	}
}

// Server handles websocket server client
func Server() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c, err := gosocketio.Dial(
		gosocketio.GetUrl("localhost", 5000, false),
		transport.GetDefaultWebsocketTransport())

	if err != nil {
		log.Fatal(err)
	}

	err = c.On(gosocketio.OnConnection, func(h *gosocketio.Channel) {
		log.Println("Connected")
	})
	if err != nil {
		log.Println("something")
		log.Fatal(err)
	}

	err = c.On(gosocketio.OnDisconnection, func(h *gosocketio.Channel) {
		log.Fatal("Disconnected")
	})
	if err != nil {
		log.Fatal(err)
	}

	// time.Sleep(1 * time.Second)

	helpers.SetInterval(func() { sendJoin(c) }, 5)

	// time.Sleep(60 * time.Second)
	// c.Close()

	// log.Println(" [x] Complete")

	serveMux := http.NewServeMux()
	serveMux.Handle("/node/", serveMux)

	log.Println("Starting server...")
	log.Panic(http.ListenAndServe(":4760", serveMux))
}
