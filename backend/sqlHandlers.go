package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)
const dsn = "user=dummy password=new_password host=localhost port=5433 dbname=hotel sslmode=disable"
func openConnection() error {
	// Replace these with your actual database connection details
	
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	defer db.Close()

	// Test the database connection
	err = db.Ping()
	if err != nil {
		return err
	}

	log.Println("Database connection is healthy.")
	return nil
}
type Room struct {
	ID int
	room 

}
//rooms  
func  postARoom(r room ) (roomSqlSchemaRsp, error){
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		 return roomSqlSchemaRsp{} ,err
	}

	defer db.Close()
	
	query := `
		INSERT INTO room (name, type,  cappicity , room_number)
		VALUES ($1, $2, $3, $4)   RETURNING id `
	stmt,err := db.Prepare(query)
	if err != nil{
		return roomSqlSchemaRsp{} ,err
	}
	var id int
	stmt.QueryRow(r.Name, r.Type ,r.Cappicity , r.RoomNumber).Scan(&id)

	if  err != nil{
		return roomSqlSchemaRsp{} ,err
	}
	
	  
    query2 := `SELECT * FROM room WHERE id = $1`
    stmt2, err := db.Prepare(query2)
    if err != nil {
        return roomSqlSchemaRsp{}, err
    }

    var room roomSqlSchema
    err = stmt2.QueryRow(id).Scan(&room.Id, &room.Name, &room.Type, &room.Capacity, &room.RoomNumber)
    if err != nil {
        return roomSqlSchemaRsp{}, err
    }



    return room.mapResponse(), nil

}

func  getRooms(  )([] roomSqlSchemaRsp , error){
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil ,err
	}

	defer db.Close()
    query := `
        SELECT id, name, type, cappicity, room_number
        FROM room`
	
    stmt , err:= db.Prepare(query)
    if err != nil {
        return nil , err
    }
	rows , err := stmt.Query()
	if err != nil {
        return nil , err
    }
    var rooms []roomSqlSchemaRsp

    for rows.Next() {
        var room roomSqlSchema
        if err := rows.Scan(&room.Id, &room.Name, &room.Type, &room.Capacity, &room.RoomNumber); err != nil {
            return nil, err
	
        }
        rooms = append(rooms, room.mapResponse())
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

	return  rooms ,nil

}
func getRoom(id int )(roomSqlSchemaRsp ,   error){
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return  roomSqlSchemaRsp{} ,err
	}

	defer db.Close()
	query := `SELECT * FROM room WHERE id = $1`
    stmt, err := db.Prepare(query)
    if err != nil {
        return roomSqlSchemaRsp{}, err
    }

    var room roomSqlSchema
    err = stmt.QueryRow(id).Scan(&room.Id, &room.Name, &room.Type, &room.Capacity, &room.RoomNumber)
    if err != nil {
        return roomSqlSchemaRsp{}, err
    }



    return room.mapResponse(), nil
}
func putRoom(id int ,r room )(roomSqlSchemaRsp ,   error){
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		 return roomSqlSchemaRsp{} ,err
	}

	defer db.Close()
	
	query := `
		UPDATE	room SET name=$1  , type=$2  ,cappicity=$3  , room_number=$4
		WHERE id=$5`
	stmt,err := db.Prepare(query)
	if err != nil{
		return roomSqlSchemaRsp{} ,err
	}

	stmt.QueryRow(r.Name, r.Type ,r.Cappicity , r.RoomNumber ,id)

	if  err != nil{
		return roomSqlSchemaRsp{} ,err
	}
	
	  
    query2 := `SELECT * FROM room WHERE id = $1`
    stmt2, err := db.Prepare(query2)
    if err != nil {
        return roomSqlSchemaRsp{}, err
    }

    var room roomSqlSchema
    err = stmt2.QueryRow(id).Scan(&room.Id, &room.Name, &room.Type, &room.Capacity, &room.RoomNumber)
    if err != nil {
        return roomSqlSchemaRsp{}, err
    }



    return room.mapResponse(), nil

}
func deleteRoom(id  int) error {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		 return err
	}

	defer db.Close()
	
	query := `DELETE FROM room WHERE id=$1`
	stmt,err := db.Prepare(query )
		if err != nil{
		return err
	}
	_,errQuery := stmt.Query(id)
	return errQuery
}
//users 
func postUser(u register) (userRspSchema , error){
	user,err:= u.createUser()
	if err != nil {
		return userRspSchema{} , err
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		 return userRspSchema{} ,err
	}

	defer db.Close()
	
	_,errUserExist := findUserByUsername(user.User , db )
	if errUserExist == nil {
		return 	userRspSchema{} ,errors.New("user name  is not  unique ") 
	}

	id,errPost:= PostUserExecutioner(user ,db)
	if errPost != nil {
		return userRspSchema{ } ,errPost
	}

	userEntity ,errFind :=findUserByid(id , db)
	if errFind != nil {
		return  userRspSchema{} , errFind
	}

	  
	return userEntity.mapResponse() , nil



}
func  findUserByidPublic(id int  ) (int, UserSchema , error){
	db, err := sql.Open("postgres", dsn)
	  
	if err != nil {
		 return http.StatusInternalServerError , UserSchema{} , err
	}
	erro := db.Ping() 
		if erro != nil {

		 return http.StatusInternalServerError , UserSchema{} , erro
	}

	defer db.Close()
	user ,err :=  findUserByid(id  ,db) 
	if err != nil {
		 return http.StatusBadRequest , UserSchema{} , err
	}
	return http.StatusOK , user , nil
	

}
func  findUserByid(id int  ,  db *sql.DB ) (UserSchema , error){
	query2 := `SELECT * FROM users WHERE id = $1`
    stmt2, err := db.Prepare(query2)
    if err != nil {
        return UserSchema{}, err
    }

    var userEntity UserSchema
    err = stmt2.QueryRow(id).Scan(&userEntity.Id, &userEntity.User, &userEntity.Name , &userEntity.hash,  &userEntity.salt, &userEntity.Role)
    if err != nil {
        return UserSchema{}, err
    }
	return userEntity , nil
	
}

func  findUserByUsername(userName string, db * sql.DB ) (UserSchema ,  error){
	query2 := `SELECT * FROM users WHERE user_name = $1`
    stmt2, err := db.Prepare(query2)
    if err != nil {
        return UserSchema{}, err
    }

    var userEntity UserSchema
    err = stmt2.QueryRow(userName).Scan(&userEntity.Id, &userEntity.User, &userEntity.Name , &userEntity.hash,  &userEntity.salt, &userEntity.Role)
	if err != nil {
        return UserSchema{}, err
    }
	return  userEntity ,nil
} 
func  PostUserExecutioner(user UserSchema , db *sql.DB)  ( int , error){
	query := `
		INSERT INTO users (name, role,  user_name , hashpw , saltpw)
		VALUES ($1, $2, $3, $4 ,$5)   RETURNING id `
	stmt,err := db.Prepare(query)
	if err != nil{
		return 0,err
	}
	var id int
	stmt.QueryRow(user.Name ,user.Role , user.User ,user.hash , user.salt).Scan(&id)

	if  err != nil{
		return  0,err
	}
	return id ,nil
} 
func loginExecutioner(u login ) (string ,int  , error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		 return "" , http.StatusBadRequest ,  err
	}

	defer db.Close()
	
	user,errUserExist := findUserByUsername(u.User , db )
	if errUserExist != nil {
		return  "" , http.StatusBadRequest ,errors.New("user not  found ") 
	}

 	pass := PassHashSalt{Hash: user.hash , Salt: user.salt}

	if  !pass.valid(u.Pass) {
			return  "" , http.StatusUnauthorized ,  errors.New("Password  do not  match")

	}
	type payload struct {
		id int 
		user string
	}
	jwt ,err := issueJwt(time.Hour * 24  ,userPayload{id:user.Id ,userName : user.User } ) 



	return jwt , http.StatusOK , nil

}


func getBookings(id ...int )(int ,[] bookingRspSchema ,error ){
	var err error
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		 return   http.StatusInternalServerError ,  nil  ,  err
	}

	defer db.Close()
	//Test connection 
	err = db.Ping() 
	if err != nil {
		 return http.StatusInternalServerError , nil  , err
	}

	query := "SELECT * FROM bookings"
	if len(id) == 1{
		query += ` WHERE  client_id=$1`
	}
  
	stmt , err :=  db.Prepare(query )
	if err != nil {
		 return http.StatusBadRequest , nil  , err
	}
    var rows *sql.Rows
	if  len(id)==1 {
		rows,err = stmt.Query(id[0]) 
	}else  {
		rows,err = stmt.Query() 
	}

	if err != nil {
		 return http.StatusBadRequest , nil  , err
	}
	var bookings  [] bookingRspSchema 


	for  rows.Next() {
		var booking bookingSchema
		err := rows.Scan(&booking.Id , &booking.CheckIn , &booking.CheckOut , &booking.ClientId , &booking.RoomId )
		if err !=nil {
			 return http.StatusBadRequest , nil , err
		}
		bookings = append(bookings , booking.mapResponse())

	}
	return http.StatusOK, bookings ,nil
}


func  postBooking(b booking  ,clientId int )(int , bookingRspSchema ,error ){
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		 return   http.StatusInternalServerError ,  bookingRspSchema{}  ,  err
	}

	defer db.Close()
	//Test connection 
	err = db.Ping() 
	if err != nil {
		 return http.StatusInternalServerError , bookingRspSchema{}  , err
	}

	query := "INSERT INTO bookings(check_in, check_out , client_id , room_id) VALUES($1,$2,$3,$4) RETURNING id"

	stmt , err :=  db.Prepare(query )
	if err != nil {
		 return http.StatusBadRequest , bookingRspSchema{}  , err
	}
	const layout = "2006/01/02"
	time1,err := time.Parse(layout ,b.CheckOut)
	if err!= nil{
		return http.StatusBadRequest , bookingRspSchema{}, errors.New(fmt.Sprintf("Could not  create date  from %s \n ", b.CheckOut))
	}
	time2,err := time.Parse(layout ,b.CheckIn)
	if err!= nil{
		return http.StatusBadRequest , bookingRspSchema{}, errors.New(fmt.Sprintf("Could not  create date  from %s \n ", b.CheckIn))
	}
	var id  int 
	err = stmt.QueryRow(time1 , time2, clientId, b.RoomId).Scan(&id ) 

	if err != nil {
		 return http.StatusBadRequest , bookingRspSchema{}  , err
	}
    
	code , booking , err := findBookingByid(id)
	return code, booking ,err

}

func  putBooking(b booking  ,clientId  , pkid int )(int , bookingRspSchema ,error ){
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		 return   http.StatusInternalServerError ,  bookingRspSchema{}  ,  err
	}

	defer db.Close()
	//Test connection 
	err = db.Ping() 
	if err != nil {
		 return http.StatusInternalServerError , bookingRspSchema{}  , err
	}

	query := "UPDATE bookings SET check_in=$1  , check_out=$2  ,  room_id=$3 WHERE id=$4 and  client_id=$5 RETURNING id"

	stmt , err :=  db.Prepare(query )
	if err != nil {
		 return http.StatusBadRequest , bookingRspSchema{}  , err
	}
	const layout = "2006/01/02"
	time1,err := time.Parse(layout ,b.CheckOut)
	if err!= nil{
		return http.StatusBadRequest , bookingRspSchema{}, errors.New(fmt.Sprintf("Could not  create date  from %s \n ", b.CheckOut))
	}
	time2,err := time.Parse(layout ,b.CheckIn)
	if err!= nil{
		return http.StatusBadRequest , bookingRspSchema{}, errors.New(fmt.Sprintf("Could not  create date  from %s \n ", b.CheckIn))
	}
	var id  int 
	err = stmt.QueryRow(time1 , time2 ,  b.RoomId , pkid , clientId).Scan(&id ) 

	if err != nil {
		 return http.StatusBadRequest , bookingRspSchema{}  , err
	}
    
	code , booking , err := findBookingByid(id)
	return code, booking ,err

}

func  findBookingByid(id int  )(int , bookingRspSchema ,error ){

db, err := sql.Open("postgres", dsn)
	if err != nil {
		 return   http.StatusInternalServerError ,  bookingRspSchema{},  err
	}

	defer db.Close()
	//Test connection 
	err = db.Ping() 
	if err != nil {
		 return http.StatusInternalServerError , bookingRspSchema{}  , err
	}


	query := "SELECT * FROM bookings WHERE id=$1"

	stmt , err :=  db.Prepare(query )
	if err != nil {
		 return http.StatusBadRequest , bookingRspSchema{}  , err
	}
	var booking   bookingSchema
	err = stmt.QueryRow(id).Scan(&booking.Id , &booking.CheckIn , &booking.CheckOut , &booking.ClientId , &booking.RoomId)

	if err != nil {
		 return http.StatusBadRequest , bookingRspSchema{}  , err
	}




	return http.StatusOK, booking.mapResponse() ,nil
}
func  findBooking(id ... int ){
	
}




func  deleteBooking(role string , pkid   , uid int  )(int  , error){
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		 return   http.StatusInternalServerError ,  err
	}

	defer db.Close()
	//Test connection 
	err = db.Ping() 
	if err != nil {
		 return http.StatusInternalServerError  , err
	}
	query := `SELECT  * FROM  bookings WHERE id=$1`
	stmt ,err :=db.Prepare(query)
	if err!= nil {
		return http.StatusBadRequest , err	
	}
	var booking  bookingRspSchema ;
	err = stmt.QueryRow(pkid).Scan(&booking.Id , &booking.CheckIn , &booking.CheckOut , &booking.ClientId , &booking.RoomId)
	if err != nil {
		return http.StatusBadRequest , err	
	}
	if role == "client"{
		if booking.ClientId != uid{
			return http.StatusForbidden , errors.New("You have  no privelleges to delete  this row ")
		}
	}
	query = `DELETE  FROM bookings WHERE id=$1`
	stmt ,err =db.Prepare(query)
	if err!= nil {
		return http.StatusBadRequest , err	
	}
	_,err = stmt.Query(pkid)
	if err != nil {
		return http.StatusBadRequest , err	
	}
	return http.StatusOK , nil
}