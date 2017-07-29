/********************************************
* Route that handles a get notes request	*
* @author: Julyamnis Marval        			*
* @version: 1.0                    			*
********************************************/

package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julymarval/example/manager"
)

// GetNotesHandler = GET = /api/notes
func GetNotesHandler(response http.ResponseWriter, requets *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	result := manager.SelectData()

	res, err := json.Marshal(result)

	if err != nil {
		log.Println("Handler - error parsing response")
	}

	response.WriteHeader(http.StatusOK)
	response.Write(res)
}
