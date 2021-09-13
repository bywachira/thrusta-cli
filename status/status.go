package status

import (
	"github.com/mackerelio/go-osstat/memory"
	"github.com/mackerelio/go-osstat/uptime"
)

type Statistics struct {
	SwapAvailable   int64
	SwapUsed        int64
	SwapTotal       int64
	Uptime          int64
	MemoryUsed      int64
	MemoryAvailable int64
	MemoryTotal     int64
}

func GetSystemStats() Statistics {
	// swap := sigar.Swap{}
	uptime, _ := uptime.Get()
	memory, _ := memory.Get()

	// swap.Get()

	var systemStatistics Statistics = Statistics{
		// SwapAvailable: swap.Free,
		// SwapTotal:     swap.Total,
		// SwapUsed:      swap.Used,
		Uptime:     uptime.Milliseconds(),
		MemoryUsed: memory.Total,
	}

	return systemStatistics
}
