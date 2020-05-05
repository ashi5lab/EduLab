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

//CreateTeacher function
func (server *Server) CreateTeacher(w http.ResponseWriter, r *http.Request) {

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

	teacher := models.Teacher{}
	err = json.Unmarshal(body, &teacher)
	fmt.Println("Teacher", teacher)
	if err != nil {
		w.WriteHeader(500)
		err := json.NewEncoder(w).Encode("Error Unmarshaling json")
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
		}
		return
	}
	err = teacher.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	createdTeacher, err := teacher.SaveTeacher(server.DB)
	if err != nil {
		json.NewEncoder(w).Encode("User not created in DB")
		return
	}
	json.NewEncoder(w).Encode(createdTeacher)
	return
}

//GetAllTeacher function
func (server *Server) GetAllTeacher(w http.ResponseWriter, r *http.Request) {

	teacher := models.Teacher{}
	teachers, err := teacher.FindAllTeachers(server.DB)
	if err != nil {
		json.NewEncoder(w).Encode("Error in getting value")
		return
	}
	json.NewEncoder(w).Encode(teachers)
}

//GetTeacher function
func (server *Server) GetTeacher(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	user := models.User{}

	if uid == uint64(user.UserID) && user.IsDeleted == true {

		responses.JSON(w, http.StatusNoContent, "")
		json.NewEncoder(w).Encode("User Not Found")

	} else {

		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		teacher := models.Teacher{}
		teacherGotten, err := teacher.FindTeacherByID(server.DB, uint32(uid))

		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		responses.JSON(w, http.StatusOK, teacherGotten)
	}

}

//UpdateTeacher function
func (server *Server) UpdateTeacher(w http.ResponseWriter, r *http.Request) {
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

	teacher := models.Teacher{}
	err = json.Unmarshal(body, &teacher)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	updatedTeacher, err := teacher.UpdateTeacher(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, updatedTeacher)
	json.NewEncoder(w).Encode("User Updated")
}

// DeleteTeacher function
func (server *Server) DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	teacher := models.Teacher{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	_, err = teacher.DeleteTeacher(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, "")
	json.NewEncoder(w).Encode("User Deleted")
}
