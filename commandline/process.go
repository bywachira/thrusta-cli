package commandline

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tesh254/thrusta-cli/helpers"
	"github.com/tesh254/thrusta-cli/services"
)

// Headers represent the default request header
type Headers struct {
	Key   string
	Value string
}

// Requests will hold all process related API calls
type Requests struct{}

// FetchActiveProcesses handles fetching all active processes
func (p *Requests) FetchActiveProcesses() Processes {
	var processes Processes

	var nodePayload helpers.NodePayload

	nodePayload = helpers.ParseJSONFile()

	var headers Headers = Headers{
		Key:   "Authorization",
		Value: nodePayload.Token,
	}

	client := &http.Client{}

	req, _ := http.NewRequest("GET", services.FormatAPIUrl("/process/active"), nil)

	req.Header.Set(headers.Key, headers.Value)

	res, _ := client.Do(req)

	json.NewDecoder(res.Body).Decode(&processes)

	return processes
}

// SendLogs handles sending of logs after each run to the API
func (p *Requests) SendLogs(logs string, nodeID string, processID string, status string) {
	client := &http.Client{}

	var nodePayload helpers.NodePayload

	nodePayload = helpers.ParseJSONFile()

	var headers Headers = Headers{
		Key:   "Authorization",
		Value: nodePayload.Token,
	}

	requestBody, _ := json.Marshal(map[string]string{
		"log":        logs,
		"process_id": processID,
		"node":       nodeID,
		"type":       status,
	})

	req, _ := http.NewRequest("POST", services.FormatAPIUrl("/process/add-log"), bytes.NewBuffer(requestBody))

	req.Header.Set("Content-Type", "application/json")

	req.Header.Set(headers.Key, headers.Value)

	client.Do(req)
}

// MakeProcessAsleep handles making a process sleep until next activation
func (p *Requests) MakeProcessAsleep(processID string) {
	client := &http.Client{}

	var nodePayload helpers.NodePayload

	nodePayload = helpers.ParseJSONFile()

	var headers Headers = Headers{
		Key:   "Authorization",
		Value: nodePayload.Token,
	}

	requestBody, _ := json.Marshal(map[string]string{
		"process_id": processID,
	})

	req, _ := http.NewRequest("PATCH", services.FormatAPIUrl("/process/sleep"), bytes.NewBuffer(requestBody))

	req.Header.Set("Content-Type", "application/json")

	req.Header.Set(headers.Key, headers.Value)

	client.Do(req)
}

// MapProcesses through all processes fetched
func (p *Requests) MapProcesses() {
	var processes Processes = p.FetchActiveProcesses()

	if processes.Count == 0 {
		log.Fatal("No active processes")
	} else {
		for _, element := range processes.Processes {
			for _, cmd := range element.Commands {
				Run(cmd.Command, p.SendLogs, element.ID)
			}
		}
	}
}

// SocketProcess handles mapping a process
func (p *Requests) SocketProcess(processes Processes) {
	if processes.Count == 0 {
		fmt.Println("No active processes")
	} else {
		for _, element := range processes.Processes {
			for _, cmd := range element.Commands {
				Run(cmd.Command, p.SendLogs, element.ID)
			}

			p.MakeProcessAsleep(element.ID)
		}
	}
}
