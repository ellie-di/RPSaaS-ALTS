package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImageString(t *testing.T) {
	v := Image{
		Type:       "type",
		CreatedBy:  "createdby",
		CreatedOn:  "createdon",
		StartedOn:  "startedon",
		FinishedOn: "finishedon",
	}

	s := v.String()

	assert.Equal(t, s, "Image:[Type:type, CreatedBy:createdby, CreatedOn:createdon, StartedOn:startedon, "+
		"FinishedOn:finishedon]")
}
