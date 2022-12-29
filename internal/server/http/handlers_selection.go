package http

import (
	"log"
	"net/http"
)

func GetSelectionHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GetSelectionHandler!")
}

func PutSelectionHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("PutSelectionHandler!")
}

func DeleteSelectionHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteSelectionHandler!")
}
