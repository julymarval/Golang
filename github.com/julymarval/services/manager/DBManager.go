/********************************************
* File that handles operations with a bd	*
* @author: Julyamnis Marval        			*
* @version: 1.0                    			*
********************************************/

package manager

import (
	"encoding/json"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/********************** STRUCTS ***********************/

/********************** CONSTS ***********************/

const (
	database = "go"
	//username   = "admin"
	//password   = "youPassword"
	collection = "users"
)

// InsertData = func to insert data in mongodb
func InsertData(data []byte) string {

	var user User

	json.Unmarshal(data, &user)

	// connect to the database
	db, err := mgo.Dial("localhost")

	if err != nil {
		log.Fatal("cannot dial mongo", err)
		response := "error"
		return response
	}

	defer db.Close() // clean up when we’re done

	col := db.DB(database).C(collection)

	// Insert Data
	err = col.Insert(user)

	if err != nil {
		log.Println(err)
		response := "error"
		return response
	}

	response := "ok"

	return response

}

// SelectData = func to select all data in a mongo collection
func SelectData() string {

	result := []User{}

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

	response := DoDataJSONResponse(0, "ok", string(res))

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
		response := DoGeneralJSONResponse(900, "internal error")
		return response
	}

	response := DoGeneralJSONResponse(0, "ok")

	return response
}

// UpdateDataById = func that updates a single field in the database
func UpdateDataById(id string, key string, value string) string {

	// connect to the database
	db, err := mgo.Dial("localhost")

	if err != nil {
		log.Fatal("cannot dial mongo", err)
		response := DoGeneralJSONResponse(900, "internal error")
		return response
	}

	defer db.Close() // clean up when we’re done

	col := db.DB(database).C(collection)

	err = col.Update(bson.M{"email": id}, bson.M{"$set": bson.M{key: value}})

	if err != nil {
		log.Printf("update fail %v\n", err)
		response := DoGeneralJSONResponse(900, "internal error")
		return response
	}

	response := DoGeneralJSONResponse(0, "ok")

	return response
}

// SelectDataByID = func that gets a doc by its ID
func SelectDataByID(id string) string {

	result := User{}

	// connect to the database
	db, err := mgo.Dial("localhost")

	if err != nil {
		log.Fatal("cannot dial mongo", err)
		return "error"
	}

	defer db.Close() // clean up when we’re done

	col := db.DB(database).C(collection)

	err = col.Find(bson.M{"email": id}).Select(bson.M{"password": 0, "_id": 0}).One(&result)

	if err != nil {
		log.Println(err)
		response := DoGeneralJSONResponse(900, "internal error")
		return response
	}

	res, err := json.Marshal(result)

	if err != nil {
		log.Println("Manager - error parsing response")
	}

	return string(res)

}

// SelectDataLogin = func that gets a doc by its ID
func SelectDataLogin(id string) string {

	result := User{}

	// connect to the database
	db, err := mgo.Dial("localhost")

	if err != nil {
		log.Fatal("cannot dial mongo", err)
		return "error"
	}

	defer db.Close() // clean up when we’re done

	col := db.DB(database).C(collection)

	err = col.Find(bson.M{"email": id}).One(&result)

	if err != nil {
		log.Println(err)
		response := DoGeneralJSONResponse(900, "internal error")
		return response
	}

	res, err := json.Marshal(result)

	if err != nil {
		log.Println("Manager - error parsing response")
	}

	return string(res)

}
