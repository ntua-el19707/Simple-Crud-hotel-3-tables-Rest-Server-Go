package main

import (
	"errors"
	"fmt"
	"time"
	"unicode"
)

type request interface {
	valid() error
}

// login request
type login struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}


func (l login) valid() error {
	return  nil

}
// register 
const  BeginCharactersLogin = 3 
const  EndCharactersLogin = 15
type register struct {
	User string `json:"user"`
	Pass string `json:"pass"`
	RePass string `json:"rePass"`
	Role  string `json:"role"`
	Name  string `json:"name"`


}
func (r register) valid() error {
	err := BetweeenSize(r.User , "user field" ,BeginCharactersLogin ,EndCharactersLogin)
	if err != nil{
		return err
	}
	err1 := PassValidator(r.Pass , "pass field" ,8)
	if err1 != nil{
		return err1
	}
	err2 := StringMatch(r.Pass , r.RePass , "pass field" , "rePass field")
	if err2 != nil{
		return err2
	}  
    err3 := BetweeenSize(r.Name , "name field" , 0,100)
		if err3!= nil{
		return err3
	}  
	err4 := validRole(r.Role)
	if err4 != nil{
		return err4
	}  
	return nil

}


/** room */ 
type room  struct {
	Name string `json:"name"`
	Type  string `json:"type"`
	Cappicity int  `json:"cappicity"`
	RoomNumber int  `json:"roomNumber"`
}


func (r room) valid() error {
	err1 := BetweeenSize(r.Name , "name  field" , 0, 100 )
	if err1 != nil{
		return err1
	}
	err2 :=  validType(r.Type)
	if err2 != nil {
		return err2
	}
	
	return  nil

}


/** booking */ 
type booking  struct   {
	
	CheckIn string `json:"checkIn"`
	CheckOut string `json:"checkOut"`

	RoomId int `json:"roomId"`


}
func (b booking) valid() error{
	if b.RoomId == 0 {
		return  errors.New(fmt.Sprintf("rooomId should be  given and  be  integer "))
	}


	err1 :=  validDate(b.CheckIn ,"checkIn field")
	if err1 != nil {
		return err1
	} 
	err2 :=  validDate(b.CheckOut ,"checkOut field")
	if err2 != nil {
		return err2
	} 
	return  nil 
} 
//* Custom  Validators

/*
BetweeenSize  - check if a strinf is between given sizes
@Param Str string
@Param  WhatIsValidating string"
@Param Begin int
@Param End  int
@returns  nil or error
*/
func BetweeenSize(Str string , WhatIsValidating string , Begin, End int) error {
	if len(Str) < Begin {
		return errors.New(fmt.Sprintf("The %s must have at  least  %d  characters " ,  WhatIsValidating , Begin))
	}
	if len(Str) > End  {
		return errors.New(fmt.Sprintf("The %s must have at  less than  %d  characters " ,  WhatIsValidating , End))
	}
	return nil
}

/*
Passvalidaoir  - check if a a password is valid
@Param Str string
@Param  WhatIsValidating string"
@Param  total int 
@returns  nil or error
*/
func PassValidator(Str string , WhatIsValidating string , total int) error {
	if len(Str) < total {
		return errors.New(fmt.Sprintf("The %s must have at  least  %d  characters " ,  WhatIsValidating ,total))
	}
	hasDigit := false
	hasUperCase := false 
	hasLowerCase := false 
	hasASymbol :=  false  

	for  _,char :=  range Str {
		if unicode.IsDigit(char){
			hasDigit = true 
		}
		if unicode.IsLower(char) {
			hasLowerCase = true 
		}
		if  unicode.IsUpper(char){
			hasUperCase =true
		}
		if !unicode.IsLetter(char) && !unicode.IsDigit(char){
			hasASymbol= true
		}
		if hasASymbol && hasDigit && hasLowerCase && hasUperCase{
			return  nil
		}
	}

	if  !hasDigit {
		return errors.New(fmt.Sprintf("The %s must have at  least  1 digit " ,  WhatIsValidating ))
	}
	if  !hasLowerCase {
		return errors.New(fmt.Sprintf("The %s must have at  least  1 lower case " ,  WhatIsValidating ))
	}
	if  !hasUperCase {
		return errors.New(fmt.Sprintf("The %s must have at  least  1 upper case " ,  WhatIsValidating ))
	}
	if  !hasASymbol {
		return errors.New(fmt.Sprintf("The %s must have at  least  1 symbol " ,  WhatIsValidating ))
	}
	return nil
}
/*
StringMatch  - check if a string matches
@Param Str1 string
@Param Str2 string
@Param  WhatIsValidating1 string
@Param  WhatIsValidating2 string

@returns  nil or error
*/
func StringMatch(Str1  , Str2 , WhatIsValidating1  , WhatIsValidating2 string) error {
	if Str1  != Str2  {
		return errors.New(fmt.Sprintf("The %s  must be equal   %s " ,  WhatIsValidating1 ,WhatIsValidating2))
	}
	return nil
}

func validInterface(i request) error{
	err := i.valid()
	return err
}
func validRole(Str string) error{
	const admin = "admin" 
	const client = "client"

	if Str != admin && Str != client{
		return errors.New(fmt.Sprintf("The role should be either %s or %s"  ,admin ,client))
	}

	return nil
}

func validType(typeOfRoom string) error{
	const standard = "standard" 
	const deluxe = "deluxe"
	const suite = "suite"

	if typeOfRoom != standard && typeOfRoom != deluxe && typeOfRoom!=suite{
		return errors.New(fmt.Sprintf("The type of room should be   either %s or %s or %s"  ,standard , deluxe ,suite))
	}

	return nil
}



func  validDate(date ,field string  ) error{
	   // Declaring layout constant
    const layout = "2006/01/02"
	_,err := time.Parse(layout , date)
	if err!= nil{
		return  errors.New(fmt.Sprintf("The %s should  have this format %s" , field , layout))
	}
	return nil
    
}