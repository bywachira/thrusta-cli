package utils

import (
	"encoding/json"
	"io/ioutil"
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
