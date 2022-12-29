package types

import (
	"fmt"
)

type Selection struct {
	CaseName          string `json:"casename"`
	CaseArea          string `json:"casearea"`
	CaseCategory      string `json:"casecategory"`
	CasePriority      string `json:"casepriority"`
	CaseTags          string `json:"casetags"`
	Action            string `json:"action"`
	Times             int    `json:"times"`
	Retry             int    `json:"retry"`
	UseNewEnvironment bool   `json:"usenewenv"`
	IgnoreFailure     bool   `json:"ignorefailure"`
}

func (selection *Selection) String() string {
	return fmt.Sprintf(
		"Selection:[CaseName:%s, CaseArea:%s, CaseCategory:%s, CasePriority:%s, CaseTags:%s,"+
			"Action:%s, Times:%d, Retry:%d, UseNewEnvironment:%t, IgnoreFailure:%t]",
		selection.CaseName, selection.CaseArea, selection.CaseCategory, selection.CasePriority, selection.CaseTags,
		selection.Action, selection.Times, selection.Retry, selection.UseNewEnvironment, selection.IgnoreFailure)
}
