package alts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJobTemplateString(t *testing.T) {
	v := JobTemplate{
		Id:         "id",
		Name:       "name",
		Tags:       "tags",
		Selections: make([]Selection, 0),
		Location:   "location",
		VmSize:     "vmsize",
		CreatedBy:  "createdby",
		CreatedOn:  "createdon",
		UpdatedBy:  "updatedby",
		UpdatedOn:  "updatedon",
	}

	s := v.String()

	assert.Equal(t, s, "JobTemplate:[Id:id, Name:name, Tags:tags, Selections:[], Location:location,"+
		"VmSize:vmsize, CreatedBy:createdby, CreatedOn:createdon, UpdatedBy:updatedby, UpdatedOn:updatedon]")
}
