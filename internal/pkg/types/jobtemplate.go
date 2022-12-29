package types

import (
	"fmt"
)

type JobTemplate struct {
	Id         string      `json:"id"`
	Name       string      `json:"name"`
	Tags       string      `json:"tags"`
	Selections []Selection `json:"selections"`
	Location   string      `json:"location"`
	VmSize     string      `json:"vmsize"`
	CreatedBy  string      `json:"createdby"`
	CreatedOn  string      `json:"createdon"`
	UpdatedBy  string      `json:"updatedby"`
	UpdatedOn  string      `json:"updatedon"`
}

func (jobTemplate *JobTemplate) String() string {
	return fmt.Sprintf(
		"JobTemplate:[Id:%s, Name:%s, Tags:%s, Selections:%v, Location:%s,"+
			"VmSize:%s, CreatedBy:%s, CreatedOn:%s, UpdatedBy:%s, UpdatedOn:%s]",
		jobTemplate.Id, jobTemplate.Name, jobTemplate.Tags, jobTemplate.Selections, jobTemplate.Location,
		jobTemplate.VmSize, jobTemplate.CreatedBy, jobTemplate.CreatedOn, jobTemplate.UpdatedBy, jobTemplate.UpdatedOn)
}
