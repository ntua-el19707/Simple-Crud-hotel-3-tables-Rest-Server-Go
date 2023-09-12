package main

import "net/http"

//messages
const ResourceNotFound = "The resource  Not Found"
/**	
	MainHandler  -  set main  handler  as  long  eith 404
	@Param w http.ResponseWriter 
	@Param  r * http.Request
*/
func MainHandler(w http.ResponseWriter  ,  r * http.Request) {
	if r.URL.Path != "/"{
		jsonErrorBuilder(w, http.StatusNotFound ,ResourceNotFound )
		return 
	}
	switch r.Method{
	default:
		jsonErrorBuilder(w ,  http.StatusMethodNotAllowed ,  methodNotImplemtent)
	}
}
/**	
	V1Handler  -  set main  handler  as  long  eith 404
	@Param w http.ResponseWriter 
	@Param  r * http.Request
*/
func V1Handler(w http.ResponseWriter  ,  r * http.Request) {
	if r.URL.Path != "/"{
		jsonErrorBuilder(w, http.StatusNotFound ,ResourceNotFound )
		return 
	}
	switch r.Method{
	default:
		jsonErrorBuilder(w ,  http.StatusMethodNotAllowed ,  methodNotImplemtent)
	}
}
/**
	setRouter -  sets  the Router for  the server 
	@Param  s * http.ServeMux -  server
*/
func setRouter(s *http.ServeMux){
	    
        s.HandleFunc("/",MainHandler)
		//version 1 	
	    v1 :=  http.NewServeMux()
		SetRouterV1(v1)
		s.Handle("/v1/",  http.StripPrefix("/v1" ,  v1))
}

//routers 
/**
	SetRouterV1 -  sets  the Router for  vesrion 1 of  the server 
	@Param  v1 * http.ServeMux -  server
*/
func SetRouterV1(v1 *http.ServeMux){
	v1.HandleFunc("/" , V1Handler)
	v1.HandleFunc("/health" , healthV1)
	v1.HandleFunc("/register",registerV1)
	v1.HandleFunc("/login" ,loginV1 )
	
	v1.HandleFunc("/rooms",roomsV1 )

	v1.HandleFunc("/bookings", bookingsV1)


}