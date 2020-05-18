package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ashi5lab/EduLab/pkg/auth"
	"github.com/ashi5lab/EduLab/pkg/models"
	"github.com/ashi5lab/EduLab/pkg/responses"
	"github.com/ashi5lab/EduLab/pkg/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

//Login method handler
func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
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

	token, userid, username, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}

	expire := time.Now().Add(20 * time.Minute)
	cookie := &http.Cookie{
		Name:     "access_token",
		Value:    token,
		HttpOnly: true,
		Expires:  expire}
	http.SetCookie(w, cookie)
	fmt.Println(cookie)
	m := models.Message{Message: "Login Success", Token: token, UserID: userid, UserName: username}

	responses.JSON(w, http.StatusOK, m)
}

//SignIn method
func (server *Server) SignIn(email, password string) (string, int, string, error) {

	var err error

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", 0, "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", 0, "", err
	}

	token, err := auth.CreateToken(user.UserID)
	if err != nil {
		return "", 0, "", err
	}
	return token, user.UserID, user.UserName, nil
}
