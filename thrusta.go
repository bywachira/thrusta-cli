// go: generate goversioninfo -icon = thrusta.ico
package main

import (
	"log"
	"os"

	"github.com/tesh254/thrusta-cli/commandline"
)

// func sendMonitor() {
// 	ticker := time.NewTicker(60 * time.Second)

// 	quit := make(chan struct{})

// 	go func() {
// 		for {
// 			select {
// 			case <-ticker.C:
// 				monitor.SendMonitorData()
// 			case <-quit:
// 				ticker.Stop()
// 				return
// 			}
// 		}
// 	}()
// }

func main() {
	app := commandline.SetupCLI()

	// authData := helpers.ParseJSONFile()

	// if len(authData.Token) > 0 {
	// 	if len(authData.Node) > 0 {
	// 		sendMonitor()
	// 	}
	// }

	err := app.Run(os.Args)

	if err != nil {
		log.Println(err)
	}
}
