package alts

import (
	"fmt"
	"log"
	"net/http"
)

type TestResult struct {
	Id            string            `json:"id"`
	JobId         string            `json:"jobid"`
	CaseName      string            `json:"casename"`
	Status        string            `json:"status"`
	Message       string            `json:"message"`
	LogUrl        string            `json:"logurl"`
	Image         string            `json:"image"`
	VmSize        string            `json:"vmsize"`
	Location      string            `json:"location"`
	KernelVersion string            `json:"kernelversion"`
	DistroVersion string            `json:"distroversion"`
	VhdGeneration string            `json:"vhdgeneration"`
	WalaVersion   string            `json:"walaversion"`
	ExtraFields   map[string]string `json:"extrafields"`
	Duration      string            `json:"duration"`
	StartTime     string            `json:"starttime"`
	EndTime       string            `json:"endtime"`
}

func (testResult *TestResult) String() string {
	return fmt.Sprintf(
		"TestResult:[Id:%s, JobId:%s, CaseName:%s, Status:%s, Message:%s, "+
			"LogUrl:%s, Image:%s, VmSize:%s, Location:%s, KernelVersion:%s, DistroVersion:%s, "+
			"VhdGeneration:%s, WalaVersion:%s, ExtraFields:%v, Duration:%s, StartTime:%s, EndTime:%s]",
		testResult.Id, testResult.JobId, testResult.CaseName, testResult.Status, testResult.Message,
		testResult.LogUrl, testResult.Image, testResult.VmSize, testResult.Location, testResult.KernelVersion,
		testResult.DistroVersion, testResult.VhdGeneration, testResult.WalaVersion, testResult.ExtraFields,
		testResult.Duration, testResult.StartTime, testResult.EndTime)
}

func GetTestResultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GetTestResultHandler!")
}

func PutTestResultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("PutTestResultHandler!")
}

func DeleteTestResultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteTestResultHandler!")
}
