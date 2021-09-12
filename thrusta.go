package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"thrusta/commandline"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	_, err := exec.Command("sudo", "chmod", "-R", "a+rwx", "/opt").Output()

	if err != nil {
		fmt.Printf("%s", err)
	}

	thrustaASCII := figure.NewFigure("Thrusta", "", true)
	thrustaASCII.Print()

	var cli commandline.CLI

	app := cli.SetupCLI()

	err = app.Run(os.Args)

	if err != nil {
		log.Println(err)
	}
}
