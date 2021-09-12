package backend

import (
	"encoding/json"
	"net/http"

	Utils "thrusta/utils"
)

type Headers struct {
	Key   string
	Value string
}

type TokenResponse struct {
	Token string `json:"token"`
}

type ErrResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type NodeBackend struct {
}

func GenerateLoginToken(apiKey string, nodeUUID string) (bool, ErrResponse) {
	backendUrl := Utils.ReadAppConfigJSON()

	var queries string = "/api/nodes/activateNode"

	var headers Headers = Headers{
		Key:   "Content-Type",
		Value: "application/json",
	}

	client := &http.Client{}

	req, _ := http.NewRequest("POST", backendUrl.Urls.API+queries, nil)

	req.Header.Set(headers.Key, headers.Value)
	req.Header.Set("T-Key", apiKey)
	req.Header.Set("T-Node-UUID", nodeUUID)

	res, _ := client.Do(req)

	var errResponse ErrResponse

	json.NewDecoder(res.Body).Decode(&errResponse)

	if errResponse.Status >= 400 {
		return false, errResponse
	} else {
		Utils.CreateCredentialsFile(apiKey, nodeUUID)

		return true, ErrResponse{
			Message: "",
			Status:  200,
		}
	}
}
