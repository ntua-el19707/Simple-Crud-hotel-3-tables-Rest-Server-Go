package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func jsonBuilder(w http.ResponseWriter ,  code  int ,payload  interface {} ){
	data,err :=  json.Marshal(payload) 
	if  err != nil {
		log.Printf("Failed to Marshal Json %v \n " ,payload) 
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Add( "Content-Type" ,"application/json")
	w.WriteHeader(code)
	w.Write(data)

} 
func jsonErrorBuilder(w http.ResponseWriter ,  code  int ,msg  string ){
 	if code > 499 {
		log.Printf("Server Internal Error:\n %s \nwith code: %d\n" , msg ,code)
 	} 
	type  ErrorStruct struct {
		Error string
	}
	payload := ErrorStruct{Error: msg}
	jsonBuilder(w ,code ,  payload)
} 