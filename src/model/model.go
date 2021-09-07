package model

import "time"

type User struct {
	//gorm.Model
	Id        int       `json:"id"`
	Fname     string    `json:"fname"`
	Lname     string    `json:"lname"`
	Email     string    `gorm:"email"`
	Password  string    `gorm:"password"`
	Dob       int       `json:"dob"`
	MobileNo  int       `json:"mobileNo"`
	SamagraId int       `json:"samagraId"`
	Otp       int       `json:"otp"`
	Address   string    `json:"address"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	LastLogin time.Time `json:"lastlogin"`
}
type Info struct {
	Message string `json:"message"`
}
