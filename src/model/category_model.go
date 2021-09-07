package model

type Category struct {
	//gorm.Model
	Name string `json:"name"`
	Tags string `json:"tags"`
}
