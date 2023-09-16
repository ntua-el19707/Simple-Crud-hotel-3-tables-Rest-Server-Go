package main

type SqlSchemas interface {
	mapResponse() interface{}
}

// room schema
type roomSqlSchemaRsp struct {
	Id         int
	Name       string
	Type       string
	Capacity   int
	RoomNumber int
}

type roomSqlSchema struct {
	roomSqlSchemaRsp
}

func (r roomSqlSchema) mapResponse() roomSqlSchemaRsp {
	room := roomSqlSchemaRsp{Id: r.Id, Name: r.Name, Type: r.Type, Capacity: r.Capacity, RoomNumber: r.RoomNumber}
	return room
}

type userRspSchema struct {
	Id   int
	User string
	Role string
	Name string
}
type UserSchema struct {
	Id   int
	User string
	Role string
	Name string
	hash string
	salt string
}

func (u UserSchema) mapResponse() userRspSchema {
	return userRspSchema{Id: u.Id, Name: u.Name, User: u.User, Role: u.Role}
}

func (u register) createUser() (UserSchema, error) {
	p := PassHashSalt{}
	err := p.genarate(u.Pass)
	if err != nil {
		return UserSchema{}, err
	}
	return UserSchema{Name: u.Name, User: u.User, Role: u.Role, hash: p.Hash, salt: p.Salt}, nil
}

type bookingSchema struct {
	Id       int
	CheckIn  string
	CheckOut string
	ClientId int
	RoomId   int
}
type bookingRspSchema struct {
	Id       int    `json:"id"`
	CheckIn  string `json:"checkIn"`
	CheckOut string `json:"checkOut"`
	ClientId int    `json:"clientId"`
	RoomId   int    `json:"roomId"`
}

func (b bookingSchema) mapResponse() bookingRspSchema {
	return bookingRspSchema{Id: b.Id, CheckIn: b.CheckIn, CheckOut: b.CheckOut, ClientId: b.ClientId, RoomId: b.RoomId}
}
