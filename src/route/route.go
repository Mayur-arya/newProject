package route

import (
	"home/vinita/mayur/newProject/src/controllers/category"
	"home/vinita/mayur/newProject/src/controllers/email"
	"home/vinita/mayur/newProject/src/controllers/forgotPassword"
	"home/vinita/mayur/newProject/src/controllers/login"
	"home/vinita/mayur/newProject/src/controllers/otp"
	"home/vinita/mayur/newProject/src/controllers/request_category"
	"home/vinita/mayur/newProject/src/controllers/sub_category"
	"home/vinita/mayur/newProject/src/controllers/user"

	"github.com/gorilla/mux"
)

var UserRegisterRoutes = func(router *mux.Router) {

	// user register
	router.HandleFunc("/user/registration", user.Add).Methods("POST")
	router.HandleFunc("/user/registration/list", user.List).Methods("GET")
	router.HandleFunc("/user/registration/listbyId/{id}", user.ListById).Methods("GET")
	router.HandleFunc("/user/registration/update/{id}", user.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/registration/delete/{id}", user.DeleteUser).Methods("DELETE")

	// for request category
	router.HandleFunc("/user/requestcategory", request_category.RequestCategory).Methods("POST")
	router.HandleFunc("/user/requestcategory/list", request_category.GetAllRequestCategory).Methods("GET")
	router.HandleFunc("/user/requestcategory/listbyId/{id}", request_category.GetRequestCategoryById).Methods("GET")
	router.HandleFunc("/user/requestcategory/update/{id}", request_category.UpdateRequestCategory).Methods("PUT")
	router.HandleFunc("/user/requestcategory/delete/{id}", request_category.DeleteRequestCategory).Methods("DELETE")

	// for category
	router.HandleFunc("/user/category", category.CategoryAdd).Methods("POST")
	router.HandleFunc("/user/category/list", category.GetAllCategory).Methods("GET")
	router.HandleFunc("/user/category/list/{id}", category.GetCategoryById).Methods("GET")
	router.HandleFunc("/user/category/update/{id}", category.UpdateCategory).Methods("PUT")
	router.HandleFunc("/user/category/delete/{id}", category.DeleteCategory).Methods("DELETE")

	// for subcategory
	router.HandleFunc("/user/subcategory", sub_category.SubCategory).Methods("POST")
	router.HandleFunc("/user/subcategory/list", sub_category.GetAllSubCategory).Methods("GET")
	router.HandleFunc("/user/subcategory/list/{id}", sub_category.GetAllSubCategoryById).Methods("GET")
	router.HandleFunc("/user/subcategory/update/{id}", sub_category.UpdateSubCategory).Methods("PUT")
	router.HandleFunc("/user/subcategory/delete/{id}", sub_category.DeleteSubCategory).Methods("DELETE")

	// for login
	router.HandleFunc("/user/login", login.Login).Methods("POST")

	// for otp
	router.HandleFunc("/user/login/otp", otp.OtpLogin).Methods("POST")
	router.HandleFunc("/user/login/otp/verify", otp.OtpVerify).Methods("POST")

	//for email
	router.HandleFunc("/user/sendmail", email.SendEmail).Methods("POST")

	// for forgot password
	router.HandleFunc("/user/forgotpassword", forgotPassword.ForgotPassword).Methods("POST")
	router.HandleFunc("/user/forgotpassword/reset", forgotPassword.ResetForgotPassword).Methods("POST")
}
