package sub_category

import (
	"encoding/json"
	"fmt"
	"home/vinita/mayur/newProject/src/config"
	"home/vinita/mayur/newProject/src/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllSubCategory(w http.ResponseWriter, r *http.Request) {
	var getlist []model.SubCategory
	db := config.ConnectDB()
	db.Find(&getlist)

	res, err := json.Marshal(getlist)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllSubCategoryById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	subcategory := &model.SubCategory{}
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	db := config.ConnectDB()
	db = db.Where("Id = ?", ID).Find(&subcategory)
	res, _ := json.Marshal(subcategory)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
