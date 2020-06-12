package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

//StudentClassMap struct
type StudentClassMap struct {
	StudentClassID int `gorm:"primary_key;AUTO_INCREMENT" json:"StudentClassID"`
	StudentID      int
	ClassID        int
	IsDeleted      bool      `gorm:"default:false" json:"-"`
	CreatedBy      int       `json:"-"`
	UpdatedBy      int       `json:"-"`
	UpdatedOn      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
}

//SaveStudentClassMap function
func (scm *StudentClassMap) SaveStudentClassMap(db *gorm.DB) (*StudentClassMap, error) {

	var err error
	err = db.Debug().Create(&scm).Error
	if err != nil {
		return &StudentClassMap{}, err
	}
	return scm, nil

}

//FindAllStudentClassMaps function
func (scm *StudentClassMap) FindAllStudentClassMaps(db *gorm.DB) (*[]StudentClassMap, error) {
	var err error
	studentclassmaps := []StudentClassMap{}
	err = db.Debug().Model(&StudentClassMap{}).Find(&studentclassmaps).Error

	if err != nil {
		return &[]StudentClassMap{}, err
	}
	return &studentclassmaps, nil

}

//FindStudentClassMapByID function
func (scm *StudentClassMap) FindStudentClassMapByID(db *gorm.DB, sid uint32) (*StudentClassMap, error) {
	var err error

	err = db.Debug().Model(&StudentClassMap{}).Where("student_id=?", sid).Take(scm).Error

	if err != nil {
		return &StudentClassMap{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &StudentClassMap{}, errors.New("Record Not Found")
	}
	return scm, nil

}

//UpdateStudentClassMap function
func (scm *StudentClassMap) UpdateStudentClassMap(db *gorm.DB, sid uint32) (*StudentClassMap, error) {

	println(sid)
	db = db.Debug().Model(&StudentClassMap{}).Where("student_id=?", sid).Take(&StudentClassMap{}).Update(&scm)

	if db.Error != nil {
		println("DB error")
		return &StudentClassMap{}, db.Error
	}

	//To display updated StudentClassMap
	// err := db.Debug().Model(&StudentClassMap{}).Where("student_id = ?", sid).Take(&sid).Error
	// if err != nil {
	// 	return &StudentClassMap{}, err
	// }
	return scm, nil

}
