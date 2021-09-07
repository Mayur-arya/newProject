package sub_category

import (
	"encoding/json"
	"fmt"
	"home/vinita/mayur/newProject/src/config"
	"home/vinita/mayur/newProject/src/model"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateSubCategory(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	updateSubCategory := &model.SubCategory{}
	json.Unmarshal(body, updateSubCategory)
	vars := mux.Vars(r)
	id := vars["id"]
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	fmt.Println(ID)
	// fetch from database user by id
	vars1 := mux.Vars(r)
	id1 := vars1["id"]
	subcategory := &model.SubCategory{}
	ID1, err1 := strconv.ParseInt(id1, 0, 0)
	if err1 != nil {
		fmt.Println("Error while parsing")
	}
	db := config.ConnectDB()
	db = db.Where("ID = ?", ID1).Find(&subcategory)
	if updateSubCategory.Name != "" {
		subcategory.Name = updateSubCategory.Name
	}
	if updateSubCategory.Tags != "" {
		subcategory.Tags = updateSubCategory.Tags
	}
	db.Save(&subcategory)
	//info := model.Info{Message: "Details Updated Successful."}
	//json.NewEncoder(w).Encode(info)
	res, _ := json.Marshal(subcategory)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
