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

//CreateStudentClassMapClassMap method
func (server *Server) CreateStudentClassMap(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	fmt.Println("Body", r.Body)

	if err != nil {
		w.WriteHeader(500)
		err := json.NewEncoder(w).Encode("Error reading studentclassmap")
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
		}
		return
	}

	studentclassmap := models.StudentClassMap{}
	err = json.Unmarshal(body, &studentclassmap)
	if err != nil {
		w.WriteHeader(500)
		err := json.NewEncoder(w).Encode("Error Unmarshaling json")
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
		}
		return
	}
	// studentclassmap.Prepare()
	mapedstudentclassmap, err := studentclassmap.SaveStudentClassMap(server.DB)
	if err != nil {
		json.NewEncoder(w).Encode("StudentClassMap not mapped in DB")
		return
	}
	json.NewEncoder(w).Encode(mapedstudentclassmap)
	return
}

//GetStudentClassMaps methd
func (server *Server) GetStudentClassMaps(w http.ResponseWriter, r *http.Request) {

	studentclassmap := models.StudentClassMap{}
	studentclassmaps, err := studentclassmap.FindAllStudentClassMaps(server.DB)
	if err != nil {
		json.NewEncoder(w).Encode("Error in getting value")
		return
	}
	json.NewEncoder(w).Encode(studentclassmaps)
}

//GetStudentClassMap method
func (server *Server) GetStudentClassMap(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	sid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	studentclassmap := models.StudentClassMap{}
	userGotten, err := studentclassmap.FindStudentClassMapByID(server.DB, uint32(sid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, userGotten)
}

//UpdateStudentClassMap Method
func (server *Server) UpdateStudentClassMap(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	studentclassmap := models.StudentClassMap{}
	err = json.Unmarshal(body, &studentclassmap)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	updatedStudentClassMap, err := studentclassmap.UpdateStudentClassMap(server.DB, uint32(sid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, updatedStudentClassMap)
	json.NewEncoder(w).Encode("StudentClassMap Updated")
}
