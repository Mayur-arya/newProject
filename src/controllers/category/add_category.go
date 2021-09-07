package category

import (
	"encoding/json"
	"fmt"
	"home/vinita/mayur/newProject/src/config"
	"home/vinita/mayur/newProject/src/model"
	"io/ioutil"
	"net/http"
)

func CategoryAdd(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add Category")
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		category := &model.Category{}
		err := json.Unmarshal(body, category)
		if err != nil {
			fmt.Println(err)
		}
		db := config.ConnectDB()
		db.AutoMigrate(&category)
		db.Create(&category)
		info := model.Info{Message: "Category added successfully"}
		json.NewEncoder(w).Encode(info)
	}
}
