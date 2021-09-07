package otp

import (
	//"crypto/rand"
	"encoding/json"
	"fmt"
	"home/vinita/mayur/newProject/src/config"
	"home/vinita/mayur/newProject/src/model"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func OtpLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Otp login ")
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		otp := &model.User{}
		err := json.Unmarshal(body, otp)
		if err != nil {
			fmt.Println(err)
		}
		if otp.MobileNo == 0 {
			info := model.Info{Message: "MobileNo Field is mandatory."}
			json.NewEncoder(w).Encode(info)
		} else {
			db := config.ConnectDB()
			result := map[string]interface{}{}
			db.Model(&model.User{}).Last(&result)
			fmt.Println(result["id"])
			min := 1000
			max := 9999
			rand.Seed(time.Now().UnixNano())
			randomNum := rand.Intn(max-min) + min

			//TData := model.User{Otp: randomNum, MobileNo: otp.MobileNo}
			db.Model(&model.User{}).Where("id = ?", result["id"]).Update("Otp", randomNum)

			info := model.Info{Message: "Otp sent on mobile number"}
			json.NewEncoder(w).Encode(info)
		}

	}
}
