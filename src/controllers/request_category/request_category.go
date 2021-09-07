package request_category

import (
	"encoding/json"
	"fmt"
	"home/vinita/mayur/newProject/src/config"
	"home/vinita/mayur/newProject/src/model"
	"io/ioutil"
	"net/http"
)

func RequestCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add Request category")
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		category := &model.RequestCategory{}
		err := json.Unmarshal(body, category)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(category.CategoryId)
		/*if category.CategoryId == "" {
			info := model.Info{Message: "CategoryId Field is mandatory."}
			json.NewEncoder(w).Encode(info)

		} else if category.NumberOfLabor == "" {
			info := model.Info{Message: "NumberOfLabor Field is mandatory."}
			json.NewEncoder(w).Encode(info)
		} else if category.StartDate == "" {
			info := model.Info{Message: "StartDate Field is mandatory."}
			json.NewEncoder(w).Encode(info)
		} else if category.EndDate == 0 {
			info := model.Info{Message: "EndDate Field is mandatory."}
			json.NewEncoder(w).Encode(info)
		} else if category.HoursPerDay == 0 {
			info := model.Info{Message: "HoursPerDay Field is mandatory."}
			json.NewEncoder(w).Encode(info)
		} else {*/
		db := config.ConnectDB()
		db.AutoMigrate(category)
		db.Create(&category)
		info := model.Info{Message: "Categories Added Successful."}
		json.NewEncoder(w).Encode(info)
		//}
	}

}
