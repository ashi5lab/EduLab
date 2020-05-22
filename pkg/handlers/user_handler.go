package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/ashi5lab/EduLab/pkg/auth"
	"github.com/ashi5lab/EduLab/pkg/models"
	"github.com/ashi5lab/EduLab/pkg/responses"
	"github.com/gorilla/mux"
)

//CreateUser method
func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	fmt.Println("Body", r.Body)

	if err != nil {
		w.WriteHeader(500)
		err := json.NewEncoder(w).Encode("Error reading user")
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
		}
		return
	}

	user := models.User{}

	err = json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(500)
		err := json.NewEncoder(w).Encode("Error Unmarshaling json")
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
		}
		return
	}

	fmt.Println(user.Password)
	user.Prepare()

	password := generatePassword()
	user.Password = password

	err = user.Validate("create")

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	createdUser, err := user.SaveUser(server.DB)
	if err != nil {
		json.NewEncoder(w).Encode("User not created in DB")
		return
	}
	json.NewEncoder(w).Encode(createdUser)
	return

}

func generatePassword() string {
	rand.Seed(time.Now().UnixNano())
	digits := "123456789"
	specials := "*@#"
	all := "ABCDEFGHJKLMNPQRSTUVWXYZ" +
		"abcdefghijkmnpqrstuvwxyz" +
		digits + specials
	length := 14
	buf := make([]byte, length)
	buf[0] = digits[rand.Intn(len(digits))]
	buf[1] = specials[rand.Intn(len(specials))]
	for i := 2; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	str := string(buf)
	return str
}

//GetUsers methd
func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {

	user := models.User{}
	users, err := user.FindAllUsers(server.DB)
	if err != nil {
		json.NewEncoder(w).Encode("Error in getting value")
		return
	}
	json.NewEncoder(w).Encode(users)

}

//GetProfile method
func (server *Server) GetProfile(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	userid, err := auth.ExtractTokenUserId(r)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = user.ValidateID(int(userid))

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	userGotten, err := user.FindUserByID(server.DB, int(userid))
	m := models.Profile{UserID: userGotten.UserID, UserName: userGotten.UserName, UserRole: userGotten.RoleID, Email: userGotten.Email, Gender: userGotten.Gender}
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, m)
}

//GetUser method
func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = user.ValidateID(int(uid))

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	userGotten, err := user.FindUserByID(server.DB, int(uid))

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, userGotten)
}

//UpdateUser Method
func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	fmt.Println(uid)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = user.ValidateID(int(uid))

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	updatedUser, err := user.UpdateUser(server.DB, int(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, updatedUser)
	json.NewEncoder(w).Encode("User Updated")
}

// DeleteUser method
func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	user := models.User{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = user.ValidateID(int(uid))

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	_, err = user.DeleteUser(server.DB, int(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, "")
	json.NewEncoder(w).Encode("User Deleted")
}
