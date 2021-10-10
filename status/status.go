package status

import (
	sigar "github.com/cloudfoundry/gosigar"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/uptime"
	"github.com/shirou/gopsutil/v3/mem"
)

type Statistics struct {
	SwapAvailable   uint64 `json:"swap_available"`
	SwapUsed        uint64 `json:"swap_used"`
	SwapTotal       uint64 `json:"swap_total"`
	Uptime          int64  `json:"uptime"`
	MemoryUsed      uint64 `json:"memory_used"`
	MemoryAvailable uint64 `json:"memory_available"`
	MemoryTotal     uint64 `json:"memory_total"`
	MemoryCached    uint64 `json:"memory_cached"`
	CPUUsed         uint64 `json:"cpu_used"`
	CPUTotal        uint64 `json:"cpu_total"`
	CPUFree         uint64 `json:"cpu_free"`
	TotalCPUS       uint64 `json:"cpu_count"`
}

func GetSystemStats() Statistics {
	swap := sigar.Swap{}
	uptime, _ := uptime.Get()
	memory, _ := mem.VirtualMemory()
	cpu, _ := cpu.Get()

	// swap.Get()

	var systemStatistics Statistics = Statistics{
		SwapAvailable:   swap.Free,
		SwapTotal:       swap.Total,
		SwapUsed:        swap.Used,
		Uptime:          uptime.Milliseconds(),
		MemoryUsed:      memory.Used,
		MemoryAvailable: memory.Free,
		MemoryTotal:     memory.Total,
		MemoryCached:    memory.Cached,
		CPUUsed:         cpu.User,
		CPUTotal:        cpu.Total,
		CPUFree:         cpu.Idle,
		TotalCPUS:       uint64(cpu.CPUCount),
	}

	return systemStatistics
}
