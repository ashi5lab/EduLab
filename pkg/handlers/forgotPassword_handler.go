package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"

	"github.com/ashi5lab/EduLab/pkg/models"
	"github.com/ashi5lab/EduLab/pkg/responses"
)

func sendMail(email string) {

	// Choose auth method and set it up
	auth := smtp.PlainAuth("mail", "nedungatfake@gmail.com", "fakenedungat@123", "smtp.mailtrap.io")

	user := models.User{}
	// Here we do it all: connect to our server, set up a message and send it

	to := []string{email}
	msg := []byte("To:" + user.UserName + "\r\n" +
		"Subject: Your current Password is \r\n" +
		"\r\n" +
		user.Password + "\r\n")
	err := smtp.SendMail("smtp.mailtrap.io:25", auth, "nedungatfake@gmail.com", to, msg)
	if err != nil {
		log.Fatal(err)

	}

}

//ForgotPassword method
func (server *Server) ForgotPassword(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", user.Email).Take(&user).Error
	if err == nil {
		// responses.ERROR(w, http.StatusUnprocessableEntity, err)
		// return
		sendMail(user.Email)
	}
	//email := strings.ToLower(r.PostFormValue("Email"))
	//uid := r.PostFormValue("UserID")

	return
}
