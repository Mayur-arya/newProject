package sub_category

import (
	"encoding/json"
	"fmt"
	"home/vinita/mayur/newProject/src/config"
	"home/vinita/mayur/newProject/src/model"
	"io/ioutil"
	"net/http"
)

func SubCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add SubCategory")
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		subcategory := &model.SubCategory{}
		err := json.Unmarshal(body, subcategory)
		if err != nil {
			fmt.Println(err)
		}
		db := config.ConnectDB()
		db.AutoMigrate(&subcategory)
		db.Create(&subcategory)
		info := model.Info{Message: "Subcategory added successfully"}
		json.NewEncoder(w).Encode(info)
	}
}
