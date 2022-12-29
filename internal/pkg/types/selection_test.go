package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectionString(t *testing.T) {
	v := Selection{
		CaseName:          "casename",
		CaseArea:          "casearea",
		CaseCategory:      "casecategory",
		CasePriority:      "casepriority",
		CaseTags:          "casetags",
		Action:            "action",
		Times:             1,
		Retry:             1,
		UseNewEnvironment: true,
		IgnoreFailure:     true,
	}

	s := v.String()

	assert.Equal(t, s, "Selection:[CaseName:casename, CaseArea:casearea, CaseCategory:casecategory, CasePriority:casepriority, CaseTags:casetags,"+
		"Action:action, Times:1, Retry:1, UseNewEnvironment:true, IgnoreFailure:true]")
}
