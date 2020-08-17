package commandline

import (
	"fmt"
	"os/exec"
	"strings"
)

// ParseCommand handles parsing a command to string array
func ParseCommand(command string) []string {
	return strings.Fields(command)
}

// Run handles calling both ParseCommand & RunProcess
func Run(order string) {
	arguments := ParseCommand(order)

	root := arguments[0]

	orderArgs := arguments[1:]

	RunProcess(root, orderArgs, false)
}

// RunProcess handles running command
func RunProcess(path string, args []string, debug bool) (out string, err error) {

	cmd := exec.Command(path, args...)

	var b []byte

	b, err = cmd.CombinedOutput()

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
