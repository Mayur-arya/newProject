package request_category

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

func UpdateRequestCategory(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	updateReqCategory := &model.RequestCategory{}
	json.Unmarshal(body, updateReqCategory)
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
	reqcategory := &model.RequestCategory{}
	ID1, err1 := strconv.ParseInt(id1, 0, 0)
	if err1 != nil {
		fmt.Println("Error while parsing")
	}
	db := config.ConnectDB()
	db = db.Where("ID = ?", ID1).Find(&reqcategory)
	if updateReqCategory.CategoryId != 0 {
		reqcategory.CategoryId = updateReqCategory.CategoryId
	}
	if updateReqCategory.NumberOfLabor != 0 {
		reqcategory.NumberOfLabor = updateReqCategory.NumberOfLabor
	}
	/*if updateReqCategory.StartDate != 0 {
		reqcategory.StartDate = updateReqCategory.StartDate
	}
	if updateReqCategory.EndDate != 0 {
		reqcategory.EndDate = updateReqCategory.EndDate
	}*/
	if updateReqCategory.HoursPerDay != 0 {
		reqcategory.HoursPerDay = updateReqCategory.HoursPerDay
	}
	if updateReqCategory.Longitude != 0 {
		reqcategory.Longitude = updateReqCategory.Longitude
	}
	if updateReqCategory.Latitude != 0 {
		reqcategory.Latitude = updateReqCategory.Latitude
	}
	db.Save(&reqcategory)
	//info := model.Info{Message: "Details Updated Successful."}
	//json.NewEncoder(w).Encode(info)
	res, _ := json.Marshal(reqcategory)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
