package http

import (
	"log"
	"net/http"
)

func GetImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GetImageHandler!")
}

func PutImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("PutImageHandler!")
}

func DeleteImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteImageHandler!")
}
