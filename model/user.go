package model

//User struct contains user data
type User struct {
	Firstname string
	Lastname string
	Address  Address
}

//Address struct contains address of a given user
type Address struct {
	City string
	State string
	Country string
}