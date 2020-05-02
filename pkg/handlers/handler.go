package handlers

import (
	"fmt"
	"log"

	"github.com/ashi5lab/EduLab/pkg/models"
	"github.com/jinzhu/gorm"

	//postgres import
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Server struct {
	DB *gorm.DB
}

//Initialize method
func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	server.DB.Debug().AutoMigrate(&models.User{}, &models.Student{}, &models.Class{}, &models.StudentClassMapping{}, &models.Role{}, &models.Teacher{}) //database migration

}
