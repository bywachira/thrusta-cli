package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type AppConfiguration struct {
	Urls URLS `json:"urls"`
}

type URLS struct {
	API string `json:"api"`
}

func ReadAppConfigJSON() AppConfiguration {
	jsonFile, err := os.Open("config.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var appConfig AppConfiguration

	json.Unmarshal(byteValue, &appConfig)

	return appConfig
}
