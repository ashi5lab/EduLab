package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ashi5lab/EduLab/pkg/models"
)

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	fmt.Println("Body", r.Body)

	if err != nil {
		w.WriteHeader(500)
		err := json.NewEncoder(w).Encode("Error creating user")
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
		err := json.NewEncoder(w).Encode("Error creating user")
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
		}
		return
	}

	user.Prepare()
	createdUser, err := user.SaveUser(server.DB)
	if err != nil {
		json.NewEncoder(w).Encode("User not created in DB")
		return
	}
	json.NewEncoder(w).Encode(createdUser)
	return
}
