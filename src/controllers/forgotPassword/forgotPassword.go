package forgotPassword

import (
	"encoding/json"
	"fmt"
	"home/vinita/mayur/newProject/src/config"
	"home/vinita/mayur/newProject/src/model"
	"io/ioutil"
	"net/http"
)

func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	forgotpass := &model.ForgotPassword{}
	err1 := json.Unmarshal(body, forgotpass)
	if err1 != nil {
		fmt.Println(err)
	}
	fmt.Println(forgotpass.Email)
	db := config.ConnectDB()
	data1 := []model.User{}
	db.Where("Email = ?", forgotpass.Email).Find(&data1)
	//db.Select("id").Where("Email = ?", forgotpass.Email).Find(&data1)
	//fmt.Println("Id->", user.Id)
	if len(data1) == 1 {
		res, _ := json.Marshal(data1)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)

		info := model.Info{Message: "Email Matched"}
		json.NewEncoder(w).Encode(info)
	} else {
		info := model.Info{Message: "Email not matched"}
		json.NewEncoder(w).Encode(info)
	}

}
