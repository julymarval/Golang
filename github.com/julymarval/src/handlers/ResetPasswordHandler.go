package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julymarval/services/manager"
)

// Password = struct
type Password struct {
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}

// ResetPasswordHandler = POST
func ResetPasswordHandler(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	var pass Password

	token := request.Header.Get("sessionToken")

	err := json.NewDecoder(request.Body).Decode(&pass)

	if err != nil {
		log.Println("Handler - error parsing response")
	}

	result, status := manager.ResetPassword(token, pass.NewPassword, pass.ConfirmPassword)

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
