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

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete user")
	vars := mux.Vars(r)
	id := vars["id"]
	user := &model.User{}
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	db := config.ConnectDB()
	db.Where("ID = ?", ID).Delete(user)
	info := model.Info{Message: "Deleted Successful."}
	json.NewEncoder(w).Encode(info)
}
