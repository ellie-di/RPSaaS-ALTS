package http

import (
	"log"
	"net/http"
)

func GetTestResultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GetTestResultHandler!")
}

func PutTestResultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("PutTestResultHandler!")
}

func DeleteTestResultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteTestResultHandler!")
}
