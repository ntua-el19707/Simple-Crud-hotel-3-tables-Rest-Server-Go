CREATE DATABASE  hotel; 
CREATE TYPE role AS ENUM ('admin', 'client');
create  table  users(
    id  serial not  null Primary Key , 
    user_name  varchar(15) not  null unique , 
    name varchar(100)  ,
    hashpw varchar(64) not  null,
    saltpw varchar(32) not  null ,
    role role not  null

    );
CREATE TYPE type AS ENUM ('standard', 'deluxe' ,'suite');
create table room(
    id  serial not  null Primary Key , 
    name varchar(100)  not  null ,
    type type not null , 
    cappicity  int  not null , 
    room_number int not null 
);
create table bookings(
    id  serial not  null Primary Key , 
    check_in  timestamp   not  null ,
    check_out  timestamp not null , 
    client_id integer not  null  References  users ON Delete no action  ON update cascade , 
    room_id integer not  null  References  room ON Delete  RESTRICT ON update cascade 
 
);

