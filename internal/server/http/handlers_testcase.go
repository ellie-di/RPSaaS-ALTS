package http

import (
	"log"
	"net/http"
)

func GetTestCaseHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GetTestCaseHandler!")
}

func PutTestCaseHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("PutTestCaseHandler!")
}

func DeleteTestCaseHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteTestCaseHandler!")
}
