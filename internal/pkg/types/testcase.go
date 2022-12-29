package types

import (
	"fmt"
)

type TestCase struct {
	Id                   string `json:"id"`
	Name                 string `json:"name"`
	Area                 string `json:"area"`
	Category             string `json:"category"`
	Tags                 string `json:"tags"`
	Priority             int    `json:"priority"`
	Estimation           string `json:"estimation"`
	Description          string `json:"description"`
	TestSuiteName        string `json:"testsuitename"`
	TestSuiteDescription string `json:"testsuitedescription"`
}

func (testCase *TestCase) String() string {
	return fmt.Sprintf(
		"TestCase:[Id:%s, Name:%s, Area:%s, Category:%s, Tags:%s,"+
			"Priority:%d, Estimation:%s, Description:%s, TestSuiteName:%s, TestSuiteDescription:%s]",
		testCase.Id, testCase.Name, testCase.Area, testCase.Category, testCase.Tags,
		testCase.Priority, testCase.Estimation, testCase.Description, testCase.TestSuiteName, testCase.TestSuiteDescription)
}
