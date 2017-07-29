/********************************************
* File that handles operations with a bd	*
* @author: Julyamnis Marval        			*
* @version: 1.0                    			*
********************************************/

package manager

import (
	"encoding/json"
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/********************** STRUCTS ***********************/

// Note = Struct
type Note struct {
	ID          bson.ObjectId `json:"_id,omitempty"bson:"_id,omitempty"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	CreatedDate time.Time     `json:"created_date"`
}

/********************** CONSTS ***********************/

const (
	hosts    = "ds026491.mongolab.com:26491"
	database = "go"
	//username   = "admin"
	//password   = "youPassword"
	collection = "notes"
)

// InsertData = func to insert data in mongodb
func InsertData(data []byte) string {

	var note Note

	json.Unmarshal(data, &note)

	// connect to the database
	db, err := mgo.Dial("localhost")

	if err != nil {
		log.Fatal("cannot dial mongo", err)
		response := doGeneralJSONResponse(900, "internal error")
		return response
	}

	defer db.Close() // clean up when we’re done

	col := db.DB(database).C(collection)

	// Insert Data
	err = col.Insert(note)

	if err != nil {
		log.Println(err)
		response := doGeneralJSONResponse(900, "internal error")
		return response
	}

	response := doGeneralJSONResponse(0, "ok")

	return response

}

// SelectData = func to select all data in a mongo collection
func SelectData() string {

	result := []Note{}

	// connect to the database
	db, err := mgo.Dial("localhost")

	if err != nil {
		log.Fatal("cannot dial mongo", err)
		return "error"
	}

	defer db.Close() // clean up when we’re done

	col := db.DB(database).C(collection)

	err = col.Find(nil).All(&result)

	if err != nil {
		log.Println(err)
		return "error"
	}

	res, err := json.Marshal(result)

	if err != nil {
		log.Println("Manager - error parsing response")
	}

	response := doDataJSONResponse(0, "ok", string(res))

	return response

}

// DeleteDataById = func that Deletes a doc in the database
func DeleteDataById(id string) string {

	// connect to the database
	db, err := mgo.Dial("localhost")

	if err != nil {
		log.Fatal("cannot dial mongo", err)
		return "error"
	}

	defer db.Close() // clean up when we’re done

	col := db.DB(database).C(collection)

	// Delete record
	err = col.Remove(bson.M{"_id": bson.ObjectIdHex(id)})

	if err != nil {
		log.Printf("remove fail %v\n", err)
		response := doGeneralJSONResponse(900, "internal error")
		return response
	}

	response := doGeneralJSONResponse(0, "ok")

	return response
}

// UpdateDataById = func that updates a single field in the database
func UpdateDataById(id string, key string, value string) string {

	// connect to the database
	db, err := mgo.Dial("localhost")

	if err != nil {
		log.Fatal("cannot dial mongo", err)
		response := doGeneralJSONResponse(900, "internal error")
		return response
	}

	defer db.Close() // clean up when we’re done

	col := db.DB(database).C(collection)

	err = col.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{key: value}})

	if err != nil {
		log.Printf("update fail %v\n", err)
		response := doGeneralJSONResponse(900, "internal error")
		return response
	}

	response := doGeneralJSONResponse(0, "ok")

	return response
}

// SelectDataByID = func that gets a doc by its ID
func SelectDataByID(id string) string {

	result := Note{}

	// connect to the database
	db, err := mgo.Dial("localhost")

	if err != nil {
		log.Fatal("cannot dial mongo", err)
		return "error"
	}

	defer db.Close() // clean up when we’re done

	col := db.DB(database).C(collection)

	err = col.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)

	if err != nil {
		log.Println(err)
		response := doGeneralJSONResponse(900, "internal error")
		return response
	}

	res, err := json.Marshal(result)

	if err != nil {
		log.Println("Manager - error parsing response")
	}

	response := doDataJSONResponse(0, "ok", string(res))

	return response

}
