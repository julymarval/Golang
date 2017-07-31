package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julymarval/services/manager"
)

//GetUserHandler = Get
func GetUserHandler(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	var token = request.Header.Get("sessionToken")

	result := manager.GetUser(token)

	bytes := json.RawMessage(result)

	res, err := bytes.MarshalJSON()

	if err != nil {
		log.Println("Handler - error parsing response")
	}

	response.WriteHeader(http.StatusOK)
	response.Write(res)

}
