package otp

import (
	"encoding/json"
	"fmt"
	"home/vinita/mayur/newProject/src/config"
	"home/vinita/mayur/newProject/src/model"
	"io/ioutil"
	"net/http"
)

func OtpVerify(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		otpVerify := &model.OtpCode{}
		err := json.Unmarshal(body, otpVerify)
		if err != nil {
			fmt.Println(err)
		}
		if otpVerify.MobileNo == 0 {
			info := model.Info{Message: "MobileNo Field is mandatory."}
			json.NewEncoder(w).Encode(info)
		} else if otpVerify.Otp == 0 {
			info := model.Info{Message: "Fill Otp send on mobile no."}
			json.NewEncoder(w).Encode(info)
		} else {
			db := config.ConnectDB()
			data1 := []model.User{}
			db.Where("mobile_no = ? AND otp = ?", otpVerify.MobileNo, otpVerify.Otp).Find(&data1)

			if len(data1) == 1 {
				info := model.Info{Message: "Otp verified Successful."}
				json.NewEncoder(w).Encode(info)
				result := map[string]interface{}{}
				db.Model(&model.User{}).Last(&result)
				db.Model(&model.User{}).Where("id = ?", result["id"]).Update("Otp", nil)
			} else {
				info := model.Info{Message: "Otp verified Failed"}
				json.NewEncoder(w).Encode(info)
			}

		}
	}
}
