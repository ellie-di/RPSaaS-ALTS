package types

import (
	"fmt"
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
