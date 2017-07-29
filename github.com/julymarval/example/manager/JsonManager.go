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

// JSONResponse = struct = json response struct
type JSONResponse struct {
	Response struct {
		Response string `json:"response"`
	}

	Error struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
}

// doGeneralJSONResponse = func that builds the general json response
func doGeneralJSONResponse(code int, msg string) string {

	var response JSONResponse

	response.Response.Response = ""
	response.Error.Code = code
	response.Error.Msg = msg

	res, err := json.Marshal(response)

	if err != nil {
		log.Println("error", err)
		return "error"
	}

	return string(res)

}

// doDataJSONResponse = func that builds a json response with data in it
func doDataJSONResponse(code int, msg string, data string) string {
	var response JSONResponse

	response.Response.Response = data
	response.Error.Code = code
	response.Error.Msg = msg

	res, err := json.Marshal(response)

	if err != nil {
		log.Println("error", err)
		return "error"
	}

	return string(res)
}
