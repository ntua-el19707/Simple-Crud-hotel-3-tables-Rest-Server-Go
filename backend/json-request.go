package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)
func jsonGetBody [ T request ] (w http.ResponseWriter  , r *http.Request) (T , error ){
 	
	d :=  json.NewDecoder(r.Body)
	var  payload T 
	err :=  d.Decode(&payload)
	if err != nil {
		jsonErrorBuilder(w ,http.StatusBadRequest , err.Error())
		return payload, err
	}
	err1 := validInterface(payload)
	if err1 != nil{
		jsonErrorBuilder(w ,http.StatusBadRequest , err1.Error())
		return payload, err1
	}

	return payload ,nil
	
} 
func getIdUrlParamParseToint(w  http.ResponseWriter , r  * http.Request )(int ,  error){
			queryParams := r.URL.Query()
			if !queryParams.Has("id"){
				jsonErrorBuilder(w  , http.StatusBadRequest , "id query params  is missing")
				return 0 ,  errors.New("id query params  is missing")
			}
			idstr  := queryParams.Get("id")
			id,err := strconv.Atoi(idstr)
			if  err!=nil{
				jsonErrorBuilder(w  , http.StatusBadRequest , "id query params  is not inetger")

			}
			return id ,err
}