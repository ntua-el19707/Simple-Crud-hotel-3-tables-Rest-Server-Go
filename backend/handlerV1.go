package main

import (
	"net/http"
)

//messages
const methodNotImplemtent = "The Method is not  Implemented"
const NotInTheBody = "The  request Body is  invalid "
/**	
	healthV1 -  Check if the server  is  up only Get implemented  
	@Param w http.ResponseWriter 
	@Param  r * http.Request
*/
func healthV1(w http.ResponseWriter  ,  r * http.Request) {
	switch r.Method {
		case http.MethodGet:
			jsonBuilder(w,http.StatusOK , struct{}{}) 
		default:
			jsonErrorBuilder(w ,http.StatusMethodNotAllowed , methodNotImplemtent )
	}

}

// Users  v1  Handlers 

/**	
	loginV1 - loggin controller
	@Param w http.ResponseWriter 
	@Param  r * http.Request
*/

func loginV1(w http.ResponseWriter  ,  r * http.Request) {
	switch r.Method {
	case http.MethodPost:

		payload ,err := jsonGetBody[login](w,r)
        if err != nil{
			return
		}
		jsonBuilder(w,http.StatusOK , payload) 
	default:
		jsonErrorBuilder(w ,http.StatusMethodNotAllowed , methodNotImplemtent )
	}
}

/**	
	registerV1 - loggin controller
	@Param w http.ResponseWriter 
	@Param  r * http.Request
*/

func registerV1(w http.ResponseWriter  ,  r * http.Request) {
	switch r.Method {
		case http.MethodPost:
			payload ,err := jsonGetBody[register](w,r)
        	if err != nil{
				return
			}
			jsonBuilder(w,http.StatusOK , payload) 
		default:
			jsonErrorBuilder(w ,http.StatusMethodNotAllowed , methodNotImplemtent )
	}
}


// rooms  

/**	
	roomsV1 - loggin controller
	@Param w http.ResponseWriter 
	@Param  r * http.Request
*/

func roomsV1(w http.ResponseWriter  ,  r * http.Request) {
	switch r.Method {
			case http.MethodPost:
			payload ,err := jsonGetBody[room](w,r)
        	if err != nil{
				return
			}

			jsonBuilder(w,http.StatusOK , payload) 
		case http.MethodPut:
			    _,err :=  getIdUrlParamParseToint(w,r)
				if err!= nil{
					return
				}
                 
				payload ,err := jsonGetBody[room](w,r)
				if err != nil{
				return
			}

			jsonBuilder(w,http.StatusOK , payload) 
		default:
			jsonErrorBuilder(w ,http.StatusMethodNotAllowed , methodNotImplemtent )
	}
}



// bookings 

/**	
	roomsV1 - loggin controller
	@Param w http.ResponseWriter 
	@Param  r * http.Request
*/

func bookingsV1(w http.ResponseWriter  ,  r * http.Request) {
	switch r.Method {
		case http.MethodPost:
			payload ,err := jsonGetBody[booking](w,r)
        	if err != nil{
				return
			}
			jsonBuilder(w,http.StatusOK , payload) 
		default:
			jsonErrorBuilder(w ,http.StatusMethodNotAllowed , methodNotImplemtent )
	}
}


