package alts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTestCaseString(t *testing.T) {
	v := TestCase{
		Id:                   "id",
		Name:                 "name",
		Area:                 "area",
		Category:             "category",
		Tags:                 "tags",
		Priority:             1,
		Estimation:           "1s",
		Description:          "abc",
		TestSuiteName:        "testsuite",
		TestSuiteDescription: "abcd",
	}

	s := v.String()

	assert.Equal(t, s, "TestCase:[Id:id, Name:name, Area:area, Category:category, Tags:tags,"+
		"Priority:1, Estimation:1s, Description:abc, TestSuiteName:testsuite, TestSuiteDescription:abcd]")
}
