package monitor

import (
	"fmt"
	"os"

	"github.com/mackerelio/go-osstat/memory"
)

func getMemory() MemoryPayload {
	memory, err := memory.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	memoryData := MemoryPayload{
		MemoryCached: memory.Cached,
		MemoryFree:   memory.Free,
		MemoryTotal:  memory.Total,
		MemoryUsed:   memory.Used,
	}

	return memoryData
}
