package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// NodePayload define payload to be sent to the API
type NodePayload struct {
	Node  string `json:"node"`
	Token string `json:"token"`
}

// SetInterval method to call function on set period
func SetInterval(handler func(), period time.Duration) {
	ticker := time.NewTicker(period * time.Second)

	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				handler()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

// ParseJSONFile handles parsing thrusta.json file
func ParseJSONFile() NodePayload {
	jsonFile, err := os.Open("thrusta.json")

	if err != nil {
		log.Fatal(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var nodePayload NodePayload

	json.Unmarshal(byteValue, &nodePayload)

	return nodePayload
}
