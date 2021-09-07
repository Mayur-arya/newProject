package user

import (
	"encoding/json"
	"fmt"
	"home/vinita/mayur/newProject/src/config"
	"home/vinita/mayur/newProject/src/model"
	"io/ioutil"
	"net/http"
)

func Add(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add User Registration")
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		//fmt.Println(body)
		user := &model.User{}
		err := json.Unmarshal(body, user)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(user.Fname)
		if user.Fname == "" {
			info := model.Info{Message: "Fname Field is mandatory."}
			json.NewEncoder(w).Encode(info)

		} else if user.Lname == "" {
			info := model.Info{Message: "Lname Field is mandatory."}
			json.NewEncoder(w).Encode(info)
		} else if user.Address == "" {
			info := model.Info{Message: "Address Field is mandatory."}
			json.NewEncoder(w).Encode(info)
		} else if user.MobileNo == 0 {
			info := model.Info{Message: "MobileNo Field is mandatory."}
			json.NewEncoder(w).Encode(info)
		} else if user.SamagraId == 0 {
			info := model.Info{Message: "SamagraId Field is mandatory."}
			json.NewEncoder(w).Encode(info)
		} else if user.Password == "" {
			info := model.Info{Message: "Password Field is mandatory."}
			json.NewEncoder(w).Encode(info)
		} else {
			db := config.ConnectDB()
			db.AutoMigrate(user)
			db.Create(&user)
			info := model.Info{Message: "Congratulation! registration Successful."}
			json.NewEncoder(w).Encode(info)
			/*res, _ := json.Marshal(user)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(res)*/
		}
	}
}
