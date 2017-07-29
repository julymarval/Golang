/********************************************
* Route that handles a put note request		*
* @author: Julyamnis Marval        			*
* @version: 1.0                    			*
********************************************/

package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/julymarval/example/manager"
)

// PutNotesHandler = PUT = /api/notes
func PutNotesHandler(response http.ResponseWriter, requets *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	var noteUpdate Note
	var key, value string
	id := mux.Vars(requets)
	k := id["id"]

	s := strings.Split(k, "=")

	err := json.NewDecoder(requets.Body).Decode(&noteUpdate)

	if err != nil {
		log.Printf("error en datos de entrada")
	}

	if noteUpdate.Description != "" {
		key = "description"
		value = noteUpdate.Description
	}
	if noteUpdate.Title != "" {
		key = "title"
		value = noteUpdate.Title
	}

	result := manager.UpdateDataById(s[1], key, value)

	res, err := json.Marshal(result)

	if err != nil {
		log.Println("Handler - error parsing json")
	}

	response.WriteHeader(http.StatusNoContent)
	response.Write(res)

}
