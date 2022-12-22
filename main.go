package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	alts "main.go/src"
)

type Config struct {
	Port int `envconfig:"PORT" default:"8080"`
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

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "handler got called!")
}

func main() {
	var conf Config
	if err := envconfig.Process("ALTS", &conf); err != nil {
		log.Fatalf("%v", err)
	}
	log.Println("[INFO] Starting ALTS server ...")

	gin.DisableConsoleColor()
	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello from ALTS!"})
	})
	router.GET("/jobtemplate", alts.GetJobTemplateHandler)
	router.GET("/jobtemplate/:name", alts.GetJobTemplateByNameHandler)
	router.GET("/jobtemplate/ResourceCreationValidate", alts.ValidateJobTemplateHandler)

	log.Printf("[INFO] Listening on port :%d", conf.Port)
	router.Run()
}
