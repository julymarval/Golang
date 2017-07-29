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

	"github.com/julymarval/example/handlers"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter().StrictSlash(false)

	add := handlers.PostNotesHandler
	get := handlers.GetNotesHandler
	put := handlers.PutNotesHandler
	delete := handlers.DeleteNotesHandler
	getID := handlers.GetNoteHandler

	route.HandleFunc("/api/notes", get).Methods("GET")
	route.HandleFunc("/api/notes", add).Methods("POST")
	route.HandleFunc("/api/notes/{id}", put).Methods("PUT")
	route.HandleFunc("/api/notes/{id}", delete).Methods("DELETE")
	route.HandleFunc("/api/notes/{id}", getID).Methods("GET")

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
