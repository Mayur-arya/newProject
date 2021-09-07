package login

import (
	"encoding/json"
	"fmt"
	"home/vinita/mayur/newProject/src/config"
	"home/vinita/mayur/newProject/src/model"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login")
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		//fmt.Println(body)
		login := &model.Login{}
		err := json.Unmarshal(body, login)

		/*user := &model.User{}
		err1 := json.Unmarshal(body, user)
		if err1 != nil {
			fmt.Println(err)
		}*/
		//fmt.Println("passord from regist ->", user.Password)
		fmt.Println("Paasword from login ->", login.Password)
		if err != nil {
			fmt.Println(err)
		}
		if login.SamagraId == 0 {
			info := model.Info{Message: "SamagraId Field is mandatory."}
			json.NewEncoder(w).Encode(info)

		} else if login.Password == "" {
			info := model.Info{Message: "Password Field is mandatory."}
			json.NewEncoder(w).Encode(info)
		} else {
			db := config.ConnectDB()
			data1 := []model.User{}

			db.Where("Password = ? AND samagra_id = ?", login.Password, login.SamagraId).Find(&data1)

			if len(data1) == 1 {
				/*res, _ := json.Marshal(data1)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(res)*/

				info := model.Info{Message: "Congratulation! login Successful."}
				json.NewEncoder(w).Encode(info)
			} else {
				info := model.Info{Message: "Login Failed"}
				json.NewEncoder(w).Encode(info)
			}
		}
	}
}
