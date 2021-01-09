package monitor

import (
	"fmt"
	"os"
	"time"

	"github.com/mackerelio/go-osstat/cpu"
)

func getCPU() CPUPayload {
	before, err := cpu.Get()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	time.Sleep(time.Duration(1) * time.Second)

	after, err := cpu.Get()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	total := float64(after.Total - before.Total)
	cpuData := CPUPayload{
		CPUUsage: float64(after.User-before.User) / total * 100,
		CPUTotal: float64(after.System-before.System) / total * 100,
		CPUIdle:  float64(after.Idle-before.Idle) / total * 100,
	}

	return cpuData
}
