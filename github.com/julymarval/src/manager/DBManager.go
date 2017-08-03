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
		return ""
	}

	defer db.Close() // clean up when we’re done

	col := db.DB(database).C(collection)

	// Insert Data
	err = col.Insert(user)

	if err != nil {
		log.Println(err)
		return ""
	}

	return "ok"

}

// SelectData = func to select all data in a mongo collection
func SelectData() string {

	result := []User{}

	// connect to the database
	db, err := mgo.Dial("localhost")

	if err != nil {
		log.Fatal("cannot dial mongo", err)
		return ""
	}

	defer db.Close() // clean up when we’re done

	col := db.DB(database).C(collection)

	err = col.Find(nil).All(&result)

	if err != nil {
		log.Println(err)
		return ""
	}

	res, err := json.Marshal(result)

	if err != nil {
		log.Println("Manager - error parsing response")
		return ""
	}

	return string(res)

}

// DeleteDataByID = func that Deletes a doc in the database
func DeleteDataByID(id string) string {

	// connect to the database
	db, err := mgo.Dial("localhost")

	if err != nil {
		log.Fatal("cannot dial mongo", err)
		return ""
	}

	defer db.Close() // clean up when we’re done

	col := db.DB(database).C(collection)

	// Delete record
	err = col.Remove(bson.M{"_id": bson.ObjectIdHex(id)})

	if err != nil {
		log.Printf("remove fail %v\n", err)
		return ""
	}

	return "ok"
}

// UpdateDataByID = func that updates a single field in the database
func UpdateDataByID(id string, key string, value string) string {

	// connect to the database
	db, err := mgo.Dial("localhost")

	if err != nil {
		log.Fatal("cannot dial mongo", err)
		return ""
	}

	defer db.Close() // clean up when we’re done

	col := db.DB(database).C(collection)

	err = col.Update(bson.M{"email": id}, bson.M{"$set": bson.M{key: value}})

	if err != nil {
		log.Printf("update fail %v\n", err)
		return ""
	}

	return "ok"
}

// SelectDataByID = func that gets a doc by its ID
func SelectDataByID(id string) string {

	result := User{}

	// connect to the database
	db, err := mgo.Dial("localhost")

	if err != nil {
		log.Fatal("cannot dial mongo", err)
		return ""
	}

	defer db.Close() // clean up when we’re done

	col := db.DB(database).C(collection)

	err = col.Find(bson.M{"email": id}).Select(bson.M{"password": 0, "_id": 0}).One(&result)

	if err != nil {
		log.Println(err)
		return ""
	}

	res, err := json.Marshal(result)

	if err != nil {
		log.Println("Manager - error parsing response")
		return ""
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
		return ""
	}

	defer db.Close() // clean up when we’re done

	col := db.DB(database).C(collection)

	err = col.Find(bson.M{"email": id}).One(&result)

	if err != nil {
		log.Println(err)
		return ""
	}

	res, err := json.Marshal(result)

	if err != nil {
		log.Println("Manager - error parsing response")
		return ""
	}

	return string(res)

}
