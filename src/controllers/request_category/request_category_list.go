package request_category

import (
	"encoding/json"
	"fmt"
	"home/vinita/mayur/newProject/src/config"
	"home/vinita/mayur/newProject/src/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllRequestCategory(w http.ResponseWriter, r *http.Request) {
	var getlist []model.RequestCategory
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

func GetRequestCategoryById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	reqcategory := &model.RequestCategory{}
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	db := config.ConnectDB()
	db = db.Where("Id = ?", ID).Find(&reqcategory)
	res, _ := json.Marshal(reqcategory)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
