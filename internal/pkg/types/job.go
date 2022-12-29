package types

import (
	"fmt"
)

type Job struct {
	Id                      string       `json:"id"`
	JobTemplateInstance     JobTemplate  `json:"jobtemplate"`
	Name                    string       `json:"name"`
	ResourceGroupNamePrefix string       `json:"resourcegroupnameprefix"`
	Image                   Image        `json:"image"`
	RepoFile                string       `json:"repofile"`
	KeepFailedEnvironment   bool         `json:"keepfailedenvironment"`
	Results                 []TestResult `json:"testresult"`
	Status                  string       `json:"status"`
	Duration                string       `json:"duration"`
	LogUrl                  string       `json:"logurl"`
}

func (job *Job) String() string {
	return fmt.Sprintf(
		"Job:[Id:%s, JobTemplateInstance:%v, Name:%s, ResourceGroupNamePrefix:%s, Image:%v,"+
			"RepoFile:%s, KeepFailedEnvironment:%t, Results:%v, Status:%s, Duration:%s, LogUrl:%s]",
		job.Id, job.JobTemplateInstance, job.Name, job.ResourceGroupNamePrefix,
		job.Image, job.RepoFile, job.KeepFailedEnvironment, job.Results, job.Status, job.Duration, job.LogUrl)
}
