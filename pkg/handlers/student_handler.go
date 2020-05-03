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

//CreateStudent method
func (server *Server) CreateStudent(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	fmt.Println("Body", r.Body)

	if err != nil {
		w.WriteHeader(500)
		err := json.NewEncoder(w).Encode("Error reading student")
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
		}
		return
	}

	student := models.Student{}
	err = json.Unmarshal(body, &student)
	if err != nil {
		w.WriteHeader(500)
		err := json.NewEncoder(w).Encode("Error Unmarshaling json")
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
		}
		return
	}
	student.Prepare()
	createdStudent, err := student.SaveStudent(server.DB)
	if err != nil {
		json.NewEncoder(w).Encode("Student not created in DB")
		return
	}
	json.NewEncoder(w).Encode(createdStudent)
	return
}

//GetStudents methd
func (server *Server) GetStudents(w http.ResponseWriter, r *http.Request) {

	student := models.Student{}
	users, err := student.FindAllStudents(server.DB)
	if err != nil {
		json.NewEncoder(w).Encode("Error in getting value")
		return
	}
	json.NewEncoder(w).Encode(users)
}

//GetStudent method
func (server *Server) GetStudent(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	student := models.Student{}
	userGotten, err := student.FindStudentByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, userGotten)
}

//UpdateStudent Method
func (server *Server) UpdateStudent(w http.ResponseWriter, r *http.Request) {
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

	student := models.Student{}
	err = json.Unmarshal(body, &student)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	updatedStudent, err := student.UpdateStudent(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, updatedStudent)
	json.NewEncoder(w).Encode("Student Updated")
}
