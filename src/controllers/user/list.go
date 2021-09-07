package user

import (
	"encoding/json"
	"fmt"
	"home/vinita/mayur/newProject/src/config"
	"home/vinita/mayur/newProject/src/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// get all list
func GetAllList() []model.User {
	db := config.ConnectDB()
	var getlist []model.User
	db.Find(&getlist)
	return getlist
}

func List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("User List Controller")
	getlist := GetAllList()
	res, err := json.Marshal(getlist)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

// get list by id
func ListById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List By Ids")
	vars := mux.Vars(r)
	id := vars["id"]
	user := &model.User{}
	Id, err := strconv.ParseInt(id, 0, 0)
	//fmt.Println(reflect.TypeOf(Id))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
	db := config.ConnectDB()
	db = db.Where("Id = ?", Id).Find(&user)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
