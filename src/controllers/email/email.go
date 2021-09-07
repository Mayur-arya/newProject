package email

import (
	"crypto/tls"
	"fmt"
	"home/vinita/mayur/newProject/src/config"
	"home/vinita/mayur/newProject/src/model"
	"net/http"

	gomail "gopkg.in/mail.v2"
)

func SendEmail(w http.ResponseWriter, r *http.Request) {
	m := gomail.NewMessage()

	db := config.ConnectDB()
	result := map[string]interface{}{}
	db.Model(&model.User{}).Last(&result)

	str := fmt.Sprintf("%v", result["email"])
	fmt.Println("emial", str)
	// Set E-Mail sender
	m.SetHeader("From", "mayurkodecreative28@gmail.com")
	// Set E-Mail receivers
	m.SetHeader("To", str)
	// Set E-Mail subject
	m.SetHeader("Subject", "NewProject test subject")
	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "mayurkodecreative28@gmail.com", "mayur@1245")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	err := d.DialAndSend(m)
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(err)

}
