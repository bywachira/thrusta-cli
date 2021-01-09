package monitor

import (
	"fmt"
	"os"

	"github.com/mackerelio/go-osstat/disk"
)

func getDisk() []DiskPayload {
	disk, err := disk.Get()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	var diskData []DiskPayload

	for i := 0; i < len(disk); i++ {
		diskData = append(diskData, DiskPayload{
			Name:   disk[i].Name,
			Reads:  disk[i].ReadsCompleted,
			Writes: disk[i].WritesCompleted,
		})
	}

	return diskData
}
