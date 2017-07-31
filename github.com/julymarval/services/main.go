/********************************************
* Main Package								*
* @author: Julyamnis Marval        			*
* @version: 1.0                    			*
********************************************/

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/julymarval/services/handlers"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter().StrictSlash(false)

	create := handlers.CreateAccountHandler
	login := handlers.LoginHandler
	get := handlers.GetUserHandler
	reset := handlers.ResetPasswordHandler
	/*update := handlers.PutNotesHandler
	 */

	route.HandleFunc("/services/createaccount", create).Methods("POST")
	route.HandleFunc("/services/login", login).Methods("POST")
	route.HandleFunc("/services/resetpassword", reset).Methods("POST")
	route.HandleFunc("/services/home", get).Methods("GET")

	server := &http.Server{
		Addr:           ":8080",
		Handler:        route,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Listening http://localhost:8080/....")

	log.Fatal(server.ListenAndServe())

}
