package model

type RequestCategory struct {
	//gorm.Model
	CategoryId    int `json:"categoryid"`
	NumberOfLabor int `json:"numberoflabor"`
	//StartDate     time.Time `json:"startdate"`
	//EndDate       time.Time `json:"enddate"`
	HoursPerDay int     `json:"hoursperday"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
}
