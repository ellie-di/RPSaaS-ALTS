package http

import (
	"log"
	"net/http"
)

func GetJobHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GetJobHandler!")
}

func PutJobHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("PutJobHandler!")
}

func DeleteJobHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteJobHandler!")
}
