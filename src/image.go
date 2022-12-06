package alts

import (
	"fmt"
	"log"
	"net/http"
)

type Image struct {
	Type       string `json:"type"`
	CreatedBy  string `json:"createdby"`
	CreatedOn  string `json:"createdon"`
	StartedOn  string `json:"startedon"`
	FinishedOn string `json:"finishedon"`
}

func (image *Image) String() string {
	return fmt.Sprintf(
		"Image:[Type:%s, CreatedBy:%s, CreatedOn:%s, StartedOn:%v, FinishedOn:%s]",
		image.Type, image.CreatedBy, image.CreatedOn, image.StartedOn, image.FinishedOn)
}

func GetImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GetImageHandler!")
}

func PutImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("PutImageHandler!")
}

func DeleteImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteImageHandler!")
}
