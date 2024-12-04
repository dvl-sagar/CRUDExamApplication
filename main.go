package main

import (
	"fmt"
	"log"
	"net/http"
)

type Configmain struct {
	MongoUri string
}

var Config Configmain

func main() {
	GetConfig()
	ConnectDB()

	mx := http.NewServeMux()
	mx.Handle("POST /student/create", http.HandlerFunc(RegisterStudent))
	mx.Handle("POST /student/get", http.HandlerFunc(GetStudent))
	mx.Handle("POST /student/update", http.HandlerFunc(UpdateStudent))
	mx.Handle("POST /student/delete", http.HandlerFunc(DeleteStudent))

	// Start the server
	fmt.Println("Server started on port 8090")
	log.Fatal(http.ListenAndServe(":8090", mx))
}
