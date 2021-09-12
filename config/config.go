package config

import (
	"fmt"
	"log"
	"os"
	"thrusta/helpers"
)

type Configuration struct {
	HasScriptsFolder bool
}

func (c *Configuration) CheckFolder(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("😈 Directory " + path + " does not exist, but not to worry we are gonna create that for you")

		c.CreateFolder(path)

		fmt.Println("✅ " + path + " was just created")
	}

	if _, err := os.Stat(path); !os.IsNotExist(err) {
		fmt.Println("✅ " + path + " already exists, noice 💯")
	}
}

func (c *Configuration) CreateFolder(path string) {
	err := os.Mkdir(path, 0777)

	var helper helpers.Helpers

	newErr := err.Error()

	if helper.ContainsString(newErr, "/opt/thrusta") {
		os.Mkdir("/opt/thrusta", 0777)
		os.Mkdir(path, 0777)
	}

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}
