package commandline

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/tesh254/thrusta/helpers"
)

// ParseCommand handles parsing a command to string array
func ParseCommand(command string) []string {
	return strings.Fields(command)
}

// Run handles calling both ParseCommand & RunProcess
func Run(order string, sendLogs func(logs string, node string, processID string, status string), processID string) {
	arguments := ParseCommand(order)

	root := arguments[0]

	orderArgs := arguments[1:]

	RunProcess(root, orderArgs, false, sendLogs, processID)
}

// RunAllProcesses handles running all processes in an array
func RunAllProcesses(processes []string, sendLogs func(logs string, node string, processID string, status string), processID string) {
	for _, element := range processes {
		Run(element, sendLogs, processID)
	}
}

// RunProcess handles running command
func RunProcess(path string, args []string, debug bool, sendLogs func(logs string, node string, processID string, status string), processID string) (out string, err error) {

	// Fields Start

	var nodePayload helpers.NodePayload

	nodePayload = helpers.ParseJSONFile()

	nodeID := nodePayload.Node

	// Fields end

	cmd := exec.Command(path, args...)

	var b []byte

	b, err = cmd.CombinedOutput()

	statStatus := "[INFO] Process has been initialized  \n"

	sendLogs(statStatus, nodeID, processID, "info")

	if err != nil {
		logError := "[ERROR] Process failed with error: \n" + string(err.Error())
		sendLogs(logError, nodeID, processID, "fail")
	}

	resultLog := "[FINAL] Process final results \n" + string(b[:len(b)])

	sendLogs(resultLog, nodeID, processID, "final")

	out = string(b)

	if debug {
		fmt.Println(strings.Join(cmd.Args[:], " "))

		if err != nil {
			fmt.Println("Run process error")
			fmt.Println(out)
		}
	}

	return
}
