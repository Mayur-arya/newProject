package forgotPassword

import (
	"encoding/json"
	"fmt"
	"home/vinita/mayur/newProject/src/config"
	"home/vinita/mayur/newProject/src/model"
	"io/ioutil"
	"net/http"
)

func ResetForgotPassword(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		fmt.Println(err)
	}
	forgotpass := &model.ForgotPassword{}
	err1 := json.Unmarshal(body, forgotpass)
	if err1 != nil {
		fmt.Println(err)
	}
	db := config.ConnectDB()
	fmt.Println(forgotpass.Password)
	fmt.Println("id->", forgotpass.Id)
	db.Model(&model.User{}).Where("Email = ?", forgotpass.Email).Update("Password", forgotpass.Password)
}
