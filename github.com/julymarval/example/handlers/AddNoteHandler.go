/********************************************
* Route that handles a post note request	*
* @author: Julyamnis Marval        			*
* @version: 1.0                    			*
********************************************/

package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/julymarval/example/manager"
)

/********************** STRUCTS ***********************/

// Note = Struct
type Note struct {
	Id          bson.ObjectId `json:"_id,omitempty"bson:"_id,omitempty"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	CreatedDate time.Time     `json:"created_date"`
}

// PostNotesHandler = POST = /api/notes
func PostNotesHandler(response http.ResponseWriter, requets *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	var notes Note

	err := json.NewDecoder(requets.Body).Decode(&notes)

	if err != nil {
		log.Println("error en datos de entrada")
	}

	notes.CreatedDate = time.Now().UTC()

	j, err := json.Marshal(notes)

	if err != nil {
		log.Println("Error parseando a json")
	}

	result := manager.InsertData(j)

	res, err := json.Marshal(result)

	if err != nil {
		log.Println("Error parseando a json")
	}

	response.WriteHeader(http.StatusCreated)
	response.Write(res)

}
