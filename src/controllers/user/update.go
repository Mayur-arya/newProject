package user

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

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update user by id")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	updateList := &model.User{}
	a := json.Unmarshal(body, updateList)
	fmt.Println(a)
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
	user := &model.User{}
	ID1, err1 := strconv.ParseInt(id1, 0, 0)
	if err1 != nil {
		fmt.Println("Error while parsing")
	}
	db := config.ConnectDB()
	db = db.Where("ID = ?", ID1).Find(&user)
	fmt.Println("name -->", updateList.Fname)
	fmt.Println("name -->", user.Fname)
	if updateList.Fname != "" {
		user.Fname = updateList.Fname
	}
	if updateList.Lname != "" {
		user.Lname = updateList.Lname
	}
	/*if updateList.Dob != nil {
		user.Dob = updateList.Dob
	}
	if updateList.MobileNo != "" {
		user.MobileNo = updateList.MobileNo
	}
	if updateList.SamagraId != "" {
		user.SamagraId = updateList.SamagraId
	}*/
	if updateList.Address != "" {
		user.Address = updateList.Address
	}
	if updateList.Email != "" {
		user.Email = updateList.Email
	}
	/*if updateList.Latitude != "" {
		user.Latitude = updateList.Latitude
	}
	if updateList.Longitude != "" {
		user.Longitude = updateList.Longitude
	}*/

	db.Save(&user)
	//info := model.Info{Message: "Details Updated Successful."}
	//json.NewEncoder(w).Encode(info)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
