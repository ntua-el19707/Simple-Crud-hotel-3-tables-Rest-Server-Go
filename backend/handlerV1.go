package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"
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
		    file,_:= readFile("public_key.pem")
			log.Println(file)
			err := openConnection()
			if err != nil{
				jsonErrorBuilder(w,http.StatusBadRequest , err.Error())
				return 
			}
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
		jwt , code , err := loginExecutioner(payload)
		if err != nil {
			jsonErrorBuilder(w ,code, err.Error() )
			return
		}
		type  response  struct {
			Jwt string  `json:"jwt"`
		}
		jsonBuilder(w,code ,response{Jwt: jwt}) 
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
			  
			user ,err := postUser(payload)
			if err != nil{
					jsonErrorBuilder(w ,http.StatusBadRequest,err.Error() )
				return
			}
			jsonBuilder(w,http.StatusOK , user) 
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
			case http.MethodGet:
			queryParams := r.URL.Query()
			if !queryParams.Has("id"){
			
				rooms, err := getRooms()
				

				if err != nil{
					jsonErrorBuilder(w, http.StatusBadRequest , err.Error())
					return 
				}
				jsonBuilder(w, http.StatusOK , rooms ) 
				
			}else{
				idstr  := queryParams.Get("id")
				id,err := strconv.Atoi(idstr)
				if err != nil {
				    jsonErrorBuilder(w , http.StatusBadRequest , "Is  is not a number")
					return 
				}
				room ,err:= getRoom(id) 
				if err != nil {
					code := http.StatusBadRequest
					if err.Error() =="sql: no rows in result set"{
						code  = http.StatusNotFound
					}
				    jsonErrorBuilder(w , code ,err.Error())
					return 
				}
				jsonBuilder(w, http.StatusOK , room) 

			}
			case http.MethodPost:
			payload ,err := jsonGetBody[room](w,r)
        	if err != nil{
				return
			}
            r, err1 := postARoom(payload)
			if err1 != nil{
				jsonErrorBuilder(w ,http.StatusBadRequest ,err1.Error())
				return  
			}
		
			jsonBuilder(w,http.StatusOK , r) 
			case http.MethodPut:
			    id,err :=  getIdUrlParamParseToint(w,r)
				if err!= nil{
					return
				}
                 
				payload ,err := jsonGetBody[room](w,r)
				if err != nil{
				return

				}
			r, err1 := putRoom(id , payload)
			if err1 != nil{
				jsonErrorBuilder(w ,http.StatusBadRequest ,err1.Error())
				return  
			}
		
			jsonBuilder(w,http.StatusOK , r) 
			case http.MethodDelete:
			    id,err :=  getIdUrlParamParseToint(w,r)
				if err!= nil{
					return
				}
                 
			
				err1 := deleteRoom(id)
			if err1 != nil{
				jsonErrorBuilder(w ,http.StatusBadRequest ,err1.Error())
				return  
			}
			type respone struct{
				Messages string `json:"msg"`
			}
		
			jsonBuilder(w,http.StatusOK ,respone{Messages: "Have  sucessfully deleted  row "  } ) 

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
	roleOfClient ,err := getRoleFromRequest(w,r)
	if err !=nil{
		return 
	} 
	client_id ,err :=  getIdFromRequest(w,  r) 
		if err != nil{
			return 
		}
	switch r.Method {
		case http.MethodGet:
			queryParams := r.URL.Query()
			if queryParams.Has("id"){
			    id,err :=  getIdUrlParamParseToint(w,r)
				if err!= nil{
					return
				}
				code , booking ,err := findBookingByid(id)

				if  err != nil {
					jsonErrorBuilder(w ,code , err.Error())
					return 
				}
				ownerShip(w ,booking , roleOfClient , client_id)
			
				
				
			}else{

			var code int 
			var bookings [] bookingRspSchema 
			var err  error 
			if roleOfClient == "client"{
				code , bookings , err = getBookings(client_id)
			}else {
					code , bookings , err = getBookings()
			}
			if err != nil {
				jsonErrorBuilder(w, code,  err.Error())
				return 
			}
			jsonBuilder(w, code, bookings)}
		case http.MethodPost:
		
			payload ,err := jsonGetBody[booking](w,r)
        	if err != nil{
				return
			}

			code , b , err := postBooking(payload , client_id)
			if err != nil {
				jsonErrorBuilder(w ,code , err.Error())
				return
			}
			jsonBuilder(w,http.StatusOK , b) 
		case http.MethodPut:
			id,err :=  getIdUrlParamParseToint(w,r)
			if err!= nil{
					return
			}
			payload ,err := jsonGetBody[booking](w,r)
        	if err != nil{
				return
			}

			code , b , err := putBooking(payload , client_id , id)
			if err != nil {
				jsonErrorBuilder(w ,code , err.Error())
				return
			}
			jsonBuilder(w,http.StatusOK , b) 
		case http.MethodDelete:
			id,err :=  getIdUrlParamParseToint(w,r)
			if err!= nil{
					return
			}
			code ,err:= deleteBooking(roleOfClient , id , client_id )

			if err != nil {
				jsonErrorBuilder(w ,code , err.Error())
				return 
			}
			type msg struct{
				Msg string  `json:"msg"`
			}
			jsonBuilder(w,code, msg{Msg: "The  recod has  been deleted"})

		default:
			jsonErrorBuilder(w ,http.StatusMethodNotAllowed , methodNotImplemtent )
	}
}


func middleware(next  http.HandlerFunc) http.HandlerFunc{
	fn :=  func(w http.ResponseWriter  ,  r * http.Request){
		log.Println("Middleware  is start")
		next.ServeHTTP(w,r)
		log.Println("End")

	}
	return http.HandlerFunc(fn)
}  

type requestAuthorized  struct{
	request * http.Request
	sub int
}  
func authMiddleware(next  http.HandlerFunc ) (http.HandlerFunc ){

	fn  := func (w http.ResponseWriter  ,  r * http.Request)  {
		r ,err := autheticator(w,r)
		if err != nil{
			return 
		}
		next.ServeHTTP(w,r)	
	}
	return http.HandlerFunc(fn)  
}
func admin(next http.HandlerFunc) http.HandlerFunc {
	
	fn  := func (w http.ResponseWriter  ,  r  *http.Request)  {
		r1 , err := autheticator(w, r)

		if err != nil{
			return 
		}

		id ,err :=  getIdFromRequest(w,  r1) 
		if err != nil{
			return 
		}

		code , user , err := findUserByidPublic(id )
		if err != nil{
			jsonErrorBuilder(w ,code  , err.Error() )
			return 
		}

		if user.Role != "admin"{
			jsonErrorBuilder(w,http.StatusForbidden , "You have not  admin priveleges")
			return  
		}

		next.ServeHTTP(w,r)
	
	
		
	}
	return http.HandlerFunc(fn)  
}

func  getIdFromRequest(w http.ResponseWriter  ,  r * http.Request) (int , error ){
	
	const err = "Sub not found in context or not an int"

 	sub, ok := r.Context().Value("sub").(int)
    if !ok {
		
        jsonErrorBuilder(w,  http.StatusInternalServerError,err)
        return 0 , errors.New(err)
    }
	return sub, nil
 
}
func  autheticator(w http.ResponseWriter  ,  r  * http.Request) ( * http.Request , error) {
		token := r.Header.Get("Authorization") 
		user , err := validate(token)
		if err != nil {
			jsonErrorBuilder(w , http.StatusUnauthorized , "User  is  unauthorized")
			return  r, err
		}
	  	 if user.ExpiresAt < time.Now().Unix() {
			jsonErrorBuilder(w , http.StatusUnauthorized , "token has expired")
			return  r, err
	  	 }
	
		ctx := context.WithValue(r.Context(), "sub", user.Id)
		r =  r.WithContext(ctx)

		
		return r , nil	
}

func profiler(next  http.HandlerFunc) http.HandlerFunc {
	fn := func (w http.ResponseWriter , r * http.Request ){
		r, err := autheticator(w , r)
		if err!= nil {
			return 
		}
		id ,err :=  getIdFromRequest(w,  r) 
		if err != nil{
			return 
		}

		code , user , err := findUserByidPublic(id )
		if err != nil{
			jsonErrorBuilder(w ,code  , err.Error() )
			return 
		}
		ctx := context.WithValue(r.Context(), "sub-role", user.Role)
		r =  r.WithContext(ctx)
	
		next.ServeHTTP(w,r)

	}

	return http.HandlerFunc(fn)
}
func  checkIfThereIs(slice [] string  ,str  string ) bool{
	for  _,item := range slice {
		if  item == str {
			return true 
		}
	}
	return  false
}

// profiler  ->  blockingMethod
func  blockingRoleMethod(next http.HandlerFunc , role string , method ...string ) http.HandlerFunc{
	fn := func(w http.ResponseWriter , r * http.Request){
		check  := checkIfThereIs(method , r.Method)
		if check{
			roleOfClient ,err := getRoleFromRequest(w,r)
			if err !=nil{
				return 
			} 
			if roleOfClient == role {
				jsonErrorBuilder(w, http.StatusForbidden , "You  have  no privilleges  to access this")
				return 
			}
			
		}
		next.ServeHTTP(w,r)
	}
	return profiler(http.HandlerFunc(fn))
}
func  getRoleFromRequest(w http.ResponseWriter  ,  r * http.Request) (string, error ){
	
	const err = "Sub not found in context or not an int"

 	role, ok := r.Context().Value("sub-role").(string)
    if !ok {
		
        jsonErrorBuilder(w,  http.StatusInternalServerError,err)
        return "" , errors.New(err)
    }
	return role, nil
 

}

func  ownerShip(w http.ResponseWriter, resp bookingRspSchema ,role  string , clientId int  )  {
	if role != "admin"{
		if (resp.ClientId != clientId){
			const errMsg ="User  has no access to see this data "
			jsonErrorBuilder(w,  http.StatusForbidden , errMsg)
			return
			
		}
	}
	jsonBuilder(w,http.StatusOK , resp)
	
}

