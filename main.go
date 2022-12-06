package main

import (
	"fmt"
	"log"
	"net/http"

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

	http.HandleFunc("/", handler)
	http.HandleFunc("/testcase/get", alts.GetTestCaseHandler)
	http.HandleFunc("/testcase/put", alts.PutTestCaseHandler)
	http.HandleFunc("/testcase/delete", alts.DeleteTestCaseHandler)
	http.HandleFunc("/testresult/get", alts.GetTestResultHandler)
	http.HandleFunc("/testresult/put", alts.PutTestResultHandler)
	http.HandleFunc("/testresult/delete", alts.DeleteTestResultHandler)
	http.HandleFunc("/selection/get", alts.GetSelectionHandler)
	http.HandleFunc("/selection/put", alts.PutSelectionHandler)
	http.HandleFunc("/selection/delete", alts.DeleteSelectionHandler)
	http.HandleFunc("/jobtemplate/get", alts.GetJobTemplateHandler)
	http.HandleFunc("/jobtemplate/put", alts.PutJobTemplateHandler)
	http.HandleFunc("/jobtemplate/delete", alts.DeleteJobTemplateHandler)
	http.HandleFunc("/job/get", alts.GetJobHandler)
	http.HandleFunc("/job/put", alts.PutJobHandler)
	http.HandleFunc("/job/delete", alts.DeleteJobHandler)
	http.HandleFunc("/image/get", alts.GetImageHandler)
	http.HandleFunc("/image/put", alts.PutImageHandler)
	http.HandleFunc("/image/delete", alts.DeleteImageHandler)
	http.HandleFunc("/OnResourceCreationCompleted", OnResourceCreationCompleted)
	http.HandleFunc("/ResourceCreationValidate", ResourceCreationValidate)
	http.HandleFunc("/ResourceCreationBegin", ResourceCreationBegin)

	log.Printf("[INFO] Listening on port :%d", conf.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), logRequest(http.DefaultServeMux))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
