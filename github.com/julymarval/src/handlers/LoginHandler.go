package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julymarval/services/manager"
)

// Credentials = struct for login
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginHandler = handler for login request
func LoginHandler(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	var cred Credentials

	err := json.NewDecoder(request.Body).Decode(&cred)

	if err != nil {
		log.Println("Handler - error parsing request")
	}

	result, status := manager.Login(cred.Username, cred.Password)

	bytes := json.RawMessage(result)

	res, err := bytes.MarshalJSON()

	if err != nil {
		log.Println("Handler - error parsing response")
	}

	if status != "ok" {
		response.WriteHeader(http.StatusInternalServerError)
	} else {
		response.WriteHeader(http.StatusOK)
	}

	response.Write(res)
}
