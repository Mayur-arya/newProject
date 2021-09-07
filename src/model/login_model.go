package model

type Login struct {
	//gorm.Model
	SamagraId int    `json:"samagraId"`
	Password  string `gorm:"password"`
}
