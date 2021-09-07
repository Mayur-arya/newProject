package main

import (
	"fmt"
	"home/vinita/mayur/newProject/src/route"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Farmers and Labours")
	r := mux.NewRouter() //init router
	route.UserRegisterRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:3000", r))
}
