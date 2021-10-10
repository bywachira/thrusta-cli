package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type CredentialsConfiguration struct {
	APIKey      string `json:"api_key"`
	Uid         string `json:"node_identifier"`
	LastCreated string `json:"last_created"`
}

func CreateCredentialsFile(token string, uid string) {
	data := &CredentialsConfiguration{
		APIKey:      token,
		Uid:         uid,
		LastCreated: time.Now().UTC().String(),
	}

	jsonData, _ := json.Marshal(data)

	ioutil.WriteFile("thrusta-credentials.json", jsonData, os.ModePerm)
}

func ReadCredentialsFile() CredentialsConfiguration {
	jsonFile, err := os.Open("thrusta-credentials.json")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully opened credentials file")

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		log.Fatal(err)
	}

	var credentials CredentialsConfiguration

	json.Unmarshal(byteValue, &credentials)

	return credentials
}
