package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"main.go/internal/pkg/db"
	"main.go/internal/pkg/types"
)

const (
	mongoDBCollectionEnvVarName = "MONGODB_COLLECTION"
)

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
func createJobTemplate(jobTemplate types.JobTemplate) {
	c := db.Connect()
	ctx := context.Background()
	defer c.Disconnect(ctx)

	var collection = os.Getenv(mongoDBCollectionEnvVarName)
	if collection == "" {
		log.Fatal("missing environment variable: ", mongoDBCollectionEnvVarName)
	}

	jobTemplateCollection := c.Database(db.DbName).Collection(collection)
	r, err := jobTemplateCollection.InsertOne(ctx, types.JobTemplate{Name: jobTemplate.Name, Tags: jobTemplate.Tags,
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
		filter = bson.D{{Key: "name", Value: name}}
	}

	conn := db.Connect()
	ctx := context.Background()
	defer conn.Disconnect(ctx)

	var collection = os.Getenv(mongoDBCollectionEnvVarName)
	if collection == "" {
		log.Fatal("missing environment variable: ", mongoDBCollectionEnvVarName)
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	jobTemplateCollection := conn.Database(db.DbName).Collection(collection)
	rs, err := jobTemplateCollection.Find(ctx, filter)
	if err != nil {
		log.Fatalf("failed to list jobtemplate(s) %v", err)
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	var jobtemplates []types.JobTemplate
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

func OnResourceCreationCompleted(w http.ResponseWriter, req *http.Request) {
	log.Println("OnResourceCreationCompleted got called!")
}

func ResourceCreationValidate(w http.ResponseWriter, req *http.Request) {
	log.Println("ResourceCreationValidate got called!")
}

func ResourceCreationBegin(w http.ResponseWriter, req *http.Request) {
	log.Println("ResourceCreationBegin got called!")
}
