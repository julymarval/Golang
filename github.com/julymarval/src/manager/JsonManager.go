/********************************************
* A json response builder					*
* @author: Julyamnis Marval        			*
* @version: 1.0                    			*
********************************************/

package manager

import (
	"encoding/json"
	"log"
)

/******************************* STRUCTS **********************************/

// JSONResponse = struct = json response struct
type JSONResponse struct {
	Response string `json:"response"`

	Error struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
}

// sessionToken = Struct
type sessionToken struct {
	SessionToken string `json:"sessionToken"`
}

/************************************* FUNCS *********************************/

// DoGeneralJSONResponse = func that builds the general json response
func DoGeneralJSONResponse(code int, msg string) string {

	var response JSONResponse

	response.Error.Code = code
	response.Error.Msg = msg

	res, err := json.Marshal(response)

	if err != nil {
		log.Println("error", err)
		return "error"
	}

	return string(res)

}

// DoDataJSONResponse = func that builds a json response with data in it
func DoDataJSONResponse(code int, msg string, data string) string {

	var response JSONResponse

	response.Response = data
	response.Error.Code = code
	response.Error.Msg = msg

	res, err := json.Marshal(response)

	if err != nil {
		log.Println("error", err)
		return "error"
	}

	return string(res)
}
