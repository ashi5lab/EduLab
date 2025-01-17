package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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
	fmt.Println("User", user)
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
	createdUser, err := user.SaveUser(server.DB)
	if err != nil {
		json.NewEncoder(w).Encode("User not created in DB")
		return
	}
	json.NewEncoder(w).Encode(createdUser)
	return
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

//GetUser method
func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	user := models.User{}
	userGotten, err := user.FindUserByID(server.DB, uint32(uid))
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

	updatedUser, err := user.UpdateUser(server.DB, uint32(uid))
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

	_, err = user.DeleteUser(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, "")
	json.NewEncoder(w).Encode("User Deleted")
}
