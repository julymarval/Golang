/************************************************
* Route that handles a get note  by id request	*
* @author: Julyamnis Marval        				*
* @version: 1.0                    				*
************************************************/

package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/julymarval/example/manager"
)

// GetNoteHandler = GET
func GetNoteHandler(response http.ResponseWriter, requets *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	id := mux.Vars(requets)
	k := id["id"]

	s := strings.Split(k, "=")

	result := manager.SelectDataByID(s[1])

	res, err := json.Marshal(result)

	if err != nil {
		log.Println("Handler - error parsing response")
	}

	response.WriteHeader(http.StatusOK)
	response.Write(res)
}
