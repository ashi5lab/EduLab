package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

//Teacher struct
type Teacher struct {
	TeacherID     uint32 `gorm:"primary_key;AUTO_INCREMENT" json:"TeacherID"`
	UserID        int    `gorm:"foreignkey:UserID;association_foreignkey:UserID"`
	Qualification string `gorm:"size:30;" json:"Qualification"`
	Subject       string `gorm:"size:20;" json:"Subject"`
	ClassID       int    `gorm:""`
	//IsDeleted     bool   `gorm:"default:false" json:"-"`
	// CreatedBy     int       `json:"-"`
	// CreatedOn     time.Time `gorm:"default:CURRENT_TIMESTAMP" `
	// UpdatedBy     int       `json:"-"`
	// UpdatedOn     time.Time `gorm:"default:CURRENT_TIMESTAMP" `
	//Users User
}

//SaveTeacher function
func (t *Teacher) SaveTeacher(db *gorm.DB) (*Teacher, error) {

	var err error
	err = db.Debug().Create(&t).Error
	if err != nil {
		return &Teacher{}, err
	}
	return t, nil

}

//FindAllTeachers function
func (t *Teacher) FindAllTeachers(db *gorm.DB) (*[]Teacher, error) {
	var err error
	teachers := []Teacher{}
	//user := User{}

	err = db.Debug(). /*Preload("Users", "is_deleted=?", false)*/ Model(&User{}).Where("is_deleted=?", false).Limit(100).Model(&Teacher{}).Error

	if err != nil {
		return &[]Teacher{}, err
	}

	return &teachers, nil
}

//FindTeacherByID function
func (t *Teacher) FindTeacherByID(db *gorm.DB, uid uint32) (*Teacher, error) {
	var err error
	//user := User{}

	err = db.Debug().Model(&User{}).Where("is_deleted=? ", false).Model(&Teacher{}).Where("user_id=? ", uid).Take(t).Error

	if err != nil {
		return &Teacher{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Teacher{}, errors.New("User Not Found")
	}

	return t, nil

}

//UpdateTeacher function
func (t *Teacher) UpdateTeacher(db *gorm.DB, uid uint32) (*Teacher, error) {

	db = db.Debug().Model(&Teacher{}).Where("user_id=?", uid).Take(&Teacher{}).Update(&t)

	if db.Error != nil {
		return &Teacher{}, db.Error
	}

	//To display updated Teacher
	err := db.Debug().Model(&Teacher{}).Where("user_id = ?", uid).Take(&t).Error
	if err != nil {
		return &Teacher{}, err
	}
	return t, nil

}

//DeleteTeacher function
func (t *Teacher) DeleteTeacher(db *gorm.DB, uid uint32) (int64, error) {

	user := User{}

	if t.UserID == user.UserID {
		db = db.Debug().Model(&User{}).Where("user_id = ?", uid).Take(&User{}).UpdateColumns(
			map[string]interface{}{
				"is_deleted": true,
			},
		)

		if db.Error != nil {
			return 0, db.Error
		}

	}

	return db.RowsAffected, nil

}
