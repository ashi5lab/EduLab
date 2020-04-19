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
	if err != nil {
		w.WriteHeader(500)
		err := json.NewEncoder(w).Encode("Error creating user")
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
		}
		return
	}
	user.Prepare()
}
