package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ashi5lab/EduLab/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

//Login method handler
func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		err := json.NewEncoder(w).Encode(`{"message":"Error Login"}`)
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
		}
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(500)
		err := json.NewEncoder(w).Encode(`{"message":"Error Login"}`)
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
		}
		return
	}
	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		w.WriteHeader(500)
		err := json.NewEncoder(w).Encode(`{"message":"Invalid User details"}`)
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
		}
		return
	}
	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode(`{"message":"Login Success,"token":"` + token + `"}`)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

//sign in method
func (server *Server) SignIn(email, password string) (string, error) {

	var err error

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.createToken(user.UserID)
}
