package commandline

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/howeyc/gopass"
	"github.com/tesh254/thrusta-cli/services"
)

// CLI holds all methods
type CLI struct{}

// LoginResponse defines login response
type LoginResponse struct {
	Token string `json:"token"`
	Node  string `json:"nodeID"`
}

// ErrorResponse defines API error response types
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// AccessCredentials defines the data types to hold access token
type AccessCredentials struct {
	Token string `json:"token"`
	Node  string `json:"node"`
}

// LoginRequest sends POST login request
func (c *CLI) LoginRequest(email string, password string) {
	fmt.Println("Initializing login...")

	if len(email) == 0 {
		log.Fatal("Please provide email")
	}

	jsonData := map[string]string{"email": email, "password": password}
	jsonValue, _ := json.Marshal(jsonData)

	request, _ := http.NewRequest("POST", services.FormatAPIUrl("/login/node"), bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Println(err)
	} else {
		responseData, _ := ioutil.ReadAll(response.Body)

		var responseObject LoginResponse

		json.Unmarshal(responseData, &responseObject)

		c.SaveToken(responseObject.Token, responseObject.Node)
	}
}

// LoginAction collects email and password and calls login request
func (c *CLI) LoginAction(email string) {
	password := ""
	fmt.Println("Enter your password: ")
	maskedpassword, _ := gopass.GetPasswdMasked()
	password = string(maskedpassword)

	c.LoginRequest(email, password)
}

// SaveToken allows saving access token and node id to keychain
func (c *CLI) SaveToken(token string, nodeid string) {
	accessCredentials := AccessCredentials{
		Token: token,
		Node:  nodeid,
	}

	file, _ := json.MarshalIndent(accessCredentials, "", "")

	_ = ioutil.WriteFile("thrusta.json", file, 0644)
}
