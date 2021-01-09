package monitor

import (
	"fmt"
	"os"

	"github.com/mackerelio/go-osstat/uptime"
)

func getUptime() UptimePayload {
	uptime, err := uptime.Get()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	uptimeData := UptimePayload{
		Uptime: uptime,
	}

	return uptimeData
}
