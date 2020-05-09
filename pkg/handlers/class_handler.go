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

//CreateClass method
func (server *Server) CreateClass(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	fmt.Println("Body", r.Body)

	if err != nil {
		w.WriteHeader(500)
		err := json.NewEncoder(w).Encode("Error reading class")
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
		}
		return
	}

	class := models.Class{}
	err = json.Unmarshal(body, &class)
	fmt.Println("Class", class)
	if err != nil {
		w.WriteHeader(500)
		err := json.NewEncoder(w).Encode("Error Unmarshaling json")
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
		}
		return
	}

	//class.Prepare()
	err = class.Validate("create")

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	createdClass, err := class.SaveClass(server.DB)
	if err != nil {
		json.NewEncoder(w).Encode("Class not created in DB")
		return
	}
	json.NewEncoder(w).Encode(createdClass)
	return
}

//GetAllClass methd
func (server *Server) GetAllClass(w http.ResponseWriter, r *http.Request) {

	class := models.Class{}
	classes, err := class.FindAllClasses(server.DB)
	if err != nil {
		json.NewEncoder(w).Encode("Error in getting value")
		return
	}
	json.NewEncoder(w).Encode(classes)
}

//GetClass method
func (server *Server) GetClass(w http.ResponseWriter, r *http.Request) {
	class := models.Class{}
	vars := mux.Vars(r)
	cid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = class.Validate("getClass")

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	classGotten, err := class.FindClassByID(server.DB, uint32(cid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, classGotten)
}

//UpdateClass Method
func (server *Server) UpdateClass(w http.ResponseWriter, r *http.Request) {
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

	class := models.Class{}
	err = json.Unmarshal(body, &class)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = class.Validate("update")

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	updatedClass, err := class.UpdateClass(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, updatedClass)
	json.NewEncoder(w).Encode("Class Updated")
}

// DeleteClass method
func (server *Server) DeleteClass(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	class := models.Class{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = class.Validate("delete")

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	_, err = class.DeleteClass(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, "")
	json.NewEncoder(w).Encode("Class Deleted")
}
