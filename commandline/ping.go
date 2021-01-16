package commandline

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/tesh254/thrusta-cli/config"
	"github.com/tesh254/thrusta-cli/helpers"
	"github.com/tesh254/thrusta-cli/monitor"
)

// InitWebsocketClient starts socket client
func InitWebsocketClient() {
	var webSoc RecConn
	cliConfig := config.ReadConfig()

	protocol := config.ProtocolChecker()

	u := url.URL{Scheme: protocol, Host: cliConfig.URL, Path: cliConfig.Path}

	helpers.Interval(func() {
		monitor.SendMonitorData()
	}, 300*time.Second)

	var req Requests

	for {
		webSoc.Dial(u.String(), nil)

		webSoc.ReadMessage()

		if len([]rune(webSoc.message)) > 0 {
			var response Processes

			json.Unmarshal([]byte(webSoc.message), &response)

			req.SocketProcess(response)
		}

	}
}
