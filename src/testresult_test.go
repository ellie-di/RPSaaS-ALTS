package alts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTestResultString(t *testing.T) {
	v := TestResult{
		Id:            "id",
		JobId:         "jobid",
		CaseName:      "casename",
		Status:        "status",
		Message:       "message",
		LogUrl:        "logurl",
		Image:         "image",
		VmSize:        "vmsize",
		Location:      "location",
		KernelVersion: "kernelversion",
		DistroVersion: "distroversion",
		VhdGeneration: "vhdgeneration",
		WalaVersion:   "walaversion",
		ExtraFields: map[string]string{
			"f": "f1",
		},
		Duration:  "duration",
		StartTime: "starttime",
		EndTime:   "endtime",
	}

	s := v.String()

	assert.Equal(t, s, "TestResult:[Id:id, JobId:jobid, CaseName:casename, Status:status, Message:message, "+
		"LogUrl:logurl, Image:image, VmSize:vmsize, Location:location, KernelVersion:kernelversion, "+
		"DistroVersion:distroversion, VhdGeneration:vhdgeneration, WalaVersion:walaversion, ExtraFields:map[f:f1], "+
		"Duration:duration, StartTime:starttime, EndTime:endtime]")
}
