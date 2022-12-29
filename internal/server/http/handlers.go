package http

import (
	"log"

	"github.com/gin-gonic/gin"
)

func GetIndexHandler(c *gin.Context) {
	log.Println("GetIndexHandler!")
}

func setHandlers(router *gin.Engine) {
	router.GET("/", GetIndexHandler)
	router.GET("/jobtemplate", GetJobTemplateHandler)
	router.GET("/jobtemplate/:name", GetJobTemplateByNameHandler)
	router.GET("/jobtemplate/ResourceCreationValidate", ValidateJobTemplateHandler)
}
