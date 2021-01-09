package monitor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/tesh254/thrusta-cli/helpers"
	"github.com/tesh254/thrusta-cli/services"
)

// CPUPayload define the cpu data
type CPUPayload struct {
	CPUUsage float64 `json:"cpu_usage"`
	CPUTotal float64 `json:"cpu_total"`
	CPUIdle  float64 `json:"cpu_idle"`
}

// MemoryPayload define the memory data
type MemoryPayload struct {
	MemoryTotal  uint64 `json:"memory_total"`
	MemoryUsed   uint64 `json:"memory_used"`
	MemoryCached uint64 `json:"memory_cached"`
	MemoryFree   uint64 `json:"memory_free"`
}

// NetworkPayload define the network data
type NetworkPayload struct {
	Name     string `json:"name"`
	Receive  uint64 `json:"receive"`
	Transmit uint64 `json:"transmit"`
}

// DiskPayload defines the disk data
type DiskPayload struct {
	Name   string `json:"name"`
	Reads  uint64 `json:"reads"`
	Writes uint64 `json:"writes"`
}

// UptimePayload define the uptime data
type UptimePayload struct {
	Uptime time.Duration `json:"uptime"`
}

// MonitorPayload define the monitor payload
type MonitorPayload struct {
	CPU     CPUPayload       `json:"cpu"`
	Memory  MemoryPayload    `json:"memory"`
	Network []NetworkPayload `json:"network"`
	Disk    []DiskPayload    `json:"disk"`
	Uptime  time.Duration    `json:"uptime"`
	Node    string           `json:"node"`
}

// Headers defines request headers
type Headers struct {
	Key   string
	Value string
}

// SendMonitorData sends monitor data
func SendMonitorData() {
	client := &http.Client{}

	var monitorPayload MonitorPayload
	nodePayload := helpers.ParseJSONFile()

	monitorPayload = MonitorPayload{
		CPU:     getCPU(),
		Memory:  getMemory(),
		Network: getNetwork(),
		Disk:    getDisk(),
		Uptime:  getUptime().Uptime,
		Node:    nodePayload.Node,
	}

	var headers Headers = Headers{
		Key:   "Authorization",
		Value: nodePayload.Token,
	}

	requestBody, _ := json.Marshal(map[string]interface{}{
		"cpu":     monitorPayload.CPU,
		"memory":  monitorPayload.Memory,
		"network": monitorPayload.Network,
		"disk":    monitorPayload.Disk,
		"uptime":  monitorPayload.Uptime,
	})

	path := "/monitoring/" + nodePayload.Node

	fmt.Println(path)

	req, _ := http.NewRequest("POST", services.FormatAPIUrl(path), bytes.NewBuffer(requestBody))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(headers.Key, headers.Value)

	client.Do(req)

	fmt.Println("Sent Server Data")
}
