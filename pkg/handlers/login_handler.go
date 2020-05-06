package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

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

	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}

	m := models.Message{Message: "Login Success", Token: token}

	json.NewEncoder(w).Encode(m)
}

//SignIn method
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

	return auth.CreateToken(uint32(user.UserID))
}
