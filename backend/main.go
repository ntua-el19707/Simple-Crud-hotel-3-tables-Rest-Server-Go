package main

import (
	"log"
	"net/http"
)

func main() {
	server_port := ":8000"
	server := &http.ServeMux{}
	log.Printf("Server is  listening on Port  %s \n" ,server_port)
	err := http.ListenAndServe(server_port , server)

	if err != nil{
		log.Fatalf("Server  has fallen due  %v\n" ,err)
	}

}