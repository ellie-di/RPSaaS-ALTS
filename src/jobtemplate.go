package alts

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	mongoDBCollectionEnvVarName = "MONGODB_COLLECTION"
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

func GetJobTemplateHandler(c *gin.Context) {
	log.Println("GetJobTemplateHandler!")
	listJobTemplate("", c)
}

func GetJobTemplateByNameHandler(c *gin.Context) {
	log.Println("GetJobTemplateByNameHandler!")
	name := c.Param("name")
	listJobTemplate(name, c)
	c.String(http.StatusOK, "GetJobTemplateByName: %s", name)
}

func PutJobTemplateHandler(c *gin.Context) {
	log.Println("PutJobTemplateHandler!")
}

func DeleteJobTemplateHandler(c *gin.Context) {
	log.Println("DeleteJobTemplateHandler!")
}

func ValidateJobTemplateHandler(c *gin.Context) {
	log.Println("ValidateJobTemplateHandler!")
	name := c.Param("name")
	cnt := listJobTemplate(name, c)
	if cnt == 0 {
		c.Writer.WriteHeader(http.StatusOK)
	} else {
		c.Writer.WriteHeader(http.StatusBadRequest)
	}
}

// creates a JobTemplate
func createJobTemplate(jobTemplate JobTemplate) {
	c := connect()
	ctx := context.Background()
	defer c.Disconnect(ctx)

	var collection = os.Getenv(mongoDBCollectionEnvVarName)
	if collection == "" {
		log.Fatal("missing environment variable: ", mongoDBCollectionEnvVarName)
	}

	jobTemplateCollection := c.Database(database).Collection(collection)
	r, err := jobTemplateCollection.InsertOne(ctx, JobTemplate{Name: jobTemplate.Name, Tags: jobTemplate.Tags,
		Selections: jobTemplate.Selections, Location: jobTemplate.Location, VmSize: jobTemplate.VmSize,
		CreatedBy: jobTemplate.CreatedBy, CreatedOn: jobTemplate.CreatedOn, UpdatedBy: jobTemplate.UpdatedBy, UpdatedOn: jobTemplate.UpdatedOn})
	if err != nil {
		log.Fatalf("failed to add a job template %v", err)
	}
	fmt.Println("added job template: ", r.InsertedID)
}

// list jobtemplates
func listJobTemplate(name string, c *gin.Context) (count int) {
	var filter interface{}
	count = 0
	if name == "" {
		filter = bson.D{}
	} else {
		filter = bson.D{{"name", name}}
	}

	conn := connect()
	ctx := context.Background()
	defer conn.Disconnect(ctx)

	var collection = os.Getenv(mongoDBCollectionEnvVarName)
	if collection == "" {
		log.Fatal("missing environment variable: ", mongoDBCollectionEnvVarName)
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	jobTemplateCollection := conn.Database(database).Collection(collection)
	rs, err := jobTemplateCollection.Find(ctx, filter)
	if err != nil {
		log.Fatalf("failed to list jobtemplate(s) %v", err)
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	var jobtemplates []JobTemplate
	err = rs.All(ctx, &jobtemplates)
	if err != nil {
		log.Fatalf("failed to list jobtemplates(s) %v", err)
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	count = len(jobtemplates)
	c.SecureJSON(http.StatusOK, jobtemplates)
	return
}
