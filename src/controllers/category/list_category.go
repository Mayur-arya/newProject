package category

import (
	"encoding/json"
	"fmt"
	"home/vinita/mayur/newProject/src/config"
	"home/vinita/mayur/newProject/src/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// get all list of category
func GetAllCategory(w http.ResponseWriter, r *http.Request) {
	var getlist []model.Category
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

// get category by id

func GetCategoryById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	category := &model.Category{}
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	db := config.ConnectDB()
	db = db.Where("Id = ?", ID).Find(&category)
	res, _ := json.Marshal(category)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
