package main

import (
	"log"
	"net/http"
	"os"
	"io"
)

func main(){
	//creating logger
	logger := log.New(os.Stdout, "[ACME] ", 0)

	//create handlers
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		logger.Printf("Handler called")
		io.WriteString(w,"Hello, world!")
	})

	//register handlers
	http.Handle("/", handler)

	//start server
	http.ListenAndServe(":8080", nil)
}
