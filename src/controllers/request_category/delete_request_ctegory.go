package request_category

import (
	"encoding/json"
	"fmt"
	"home/vinita/mayur/newProject/src/config"
	"home/vinita/mayur/newProject/src/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteRequestCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	reqcategory := &model.RequestCategory{}
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	db := config.ConnectDB()
	db.Where("ID = ?", ID).Delete(reqcategory)
	info := model.Info{Message: "Deleted Successful."}
	json.NewEncoder(w).Encode(info)
}
