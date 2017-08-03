/********************************************
* Route that handles a create account		*
* @author: Julyamnis Marval        			*
* @version: 1.0                    			*
********************************************/

package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/julymarval/services/manager"
)

/********************** STRUCTS ***********************/

// User = Struct
type User struct {
	ID          bson.ObjectId `json:"_id,omitempty"bson:"_id,omitempty"`
	Name        string        `json:"name"`
	LastName    string        `json:"lastName"`
	Email       string        `json:"email"`
	DateofBirth string        `json:"dateOfBirth"`
	Password    string        `json:"password"`
}

// CreateAccountHandler = POST
func CreateAccountHandler(response http.ResponseWriter, requets *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	var user User

	err := json.NewDecoder(requets.Body).Decode(&user)

	if err != nil {
		log.Println("error en datos de entrada")
	}

	j, err := json.Marshal(user)

	if err != nil {
		log.Println("Error parseando a json")
	}

	result, status := manager.CreateAccount(j)

	bytes := json.RawMessage(result)

	res, err := bytes.MarshalJSON()

	if err != nil {
		log.Println("Error parseando a json")
	}

	if status != "ok" {
		response.WriteHeader(http.StatusInternalServerError)
	} else {
		response.WriteHeader(http.StatusCreated)
	}

	response.Write(res)

}
