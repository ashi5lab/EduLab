package models

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

//Student struct
type Student struct {
	StudentID                 int `gorm:"primary_key;AUTO_INCREMENT" json:"StudentID"`
	UserID                    int `gorm:"not null;"`
	StudentAdmno              int `gorm:"not null;unique" json:"StudentAdmno"`
	StudentSlno               int
	StudentAppno              string
	StudentGuardianName       string    `gorm:"size:50;" json:"StudentGuardianName"`
	StudentGuardianOccupation string    `gorm:"size:50;" json:"StudentGuardianOccupation"`
	StudentGuardianRelation   string    `gorm:"size:50;" json:"StudentGuardianRelation"`
	StudentAddress            string    `gorm:"size:100;" json:"StudentAddress"`
	StudentReligion           string    `gorm:"size:20;" json:"StudentReligion"`
	StudentCaste              string    `gorm:"size:20;" json:"StudentCaste"`
	StudentCategory           string    `gorm:"size:20;" json:"StudentCategory"`
	StudentIsOEC              bool      `json:"StudentIsOEC"`
	StudentIsLingMin          bool      `json:"StudentIsLingMin"`
	StudentLingMin            string    `gorm:"size:50;" json:"StudentLingMinDesc"`
	StudentDOA                time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"StudentDOA"`
	StudentAdmissionCategory  string    `gorm:"size:50;" json:"StudentAdmissionCategory"`
	IsDeleted                 bool      `gorm:"default:false"`
	CreatedBy                 int       `json:"-"`
	CreatedOn                 time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedBy                 int       `json:"-"`
	UpdatedOn                 time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
	Users                     User      `gorm:"foreignkey:UserID;association_foreignkey:UserID"`
}

//Prepare function
func (s *Student) Prepare() {
	// u.StudentName = html.EscapeString(strings.TrimSpace(u.StudentName))
	// u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	// hashedPassword, _ := Hash(u.Password)
	// u.Password = string(hashedPassword)
	// u.CreatedOn = time.Now()
	// u.UpdatedOn = time.Now()
}

//Validate method
func (s *Student) Validate(action string) error {

	switch strings.ToLower(action) {
	case "create":
		if s.UserID == 0 {
			return errors.New("Required UserID")
		}

		if s.StudentAdmno == 0 {
			return errors.New("Required StudentAdmno")
		}
		// if s.StudentSlno == 0 {
		// 	return errors.New("Required StudentSlno")
		// }

		// if s.StudentAppno == 0 {
		// 	return errors.New("Required StudentAppno")
		// }

		// if s.StudentGuardianName == "" {
		// 	return errors.New("Required StudentGuardianName")
		// }
		// if s.StudentGuardianOccupation == "" {
		// 	return errors.New("Required StudentGuardianOccupation")
		// }
		// if s.StudentGuardianRelation == "" {
		// 	return errors.New("Required StudentGuardianRelation")
		// }
		// if s.StudentAddress == "" {
		// 	return errors.New("Required StudentAddress")
		// }
		// if s.StudentReligion == "" {
		// 	return errors.New("Required StudentReligion")
		// }
		// if s.StudentCaste == "" {
		// 	return errors.New("Required StudentCaste")
		// }
		// if s.StudentCategory == "" {
		// 	return errors.New("Required StudentCategory")
		// }
		// if s.StudentCaste == "" {
		// 	return errors.New("Required StudentCaste")
		// }
		// if s.StudentCaste == "" {
		// 	return errors.New("Required StudentCaste")
		// }
		// if s.StudentCaste == "" {
		// 	return errors.New("Required StudentCaste")
		// }
		return nil

	default:
		if s.UserID == 0 {
			return errors.New("Required UserID")
		}
		return nil

	}
}

//ValidateID method
func (s *Student) ValidateID(id int) error {

	if id == 0 {
		return errors.New("Required UserID")

	}
	return nil
}

//SaveStudent method
func (s *Student) SaveStudent(db *gorm.DB) (*Student, error) {

	var err error
	err = db.Debug().Create(&s).Error
	if err != nil {
		return &Student{}, err
	}
	return s, nil

}

//FindAllStudents method
func (s *Student) FindAllStudents(db *gorm.DB) (*[]Student, error) {
	var err error
	students := []Student{}
	err = db.Debug().Preload("Users", "is_deleted=?", false).Preload("Users.Roles").Model(&Student{}).Find(&students).Error

	if err != nil {
		return &[]Student{}, err
	}
	return &students, nil

}

//FindStudentByID method
func (s *Student) FindStudentByID(db *gorm.DB, uid uint32) (*Student, error) {
	var err error

	err = db.Debug().Model(&Student{}).Where("user_id=?", uid).Take(s).Error

	if err != nil {
		return &Student{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Student{}, errors.New("Student Not Found")
	}
	return s, nil

}

//UpdateStudent a student
func (s *Student) UpdateStudent(db *gorm.DB, uid uint32) (*Student, error) {
	s.Prepare()
	db = db.Debug().Model(&Student{}).Where("user_id = ?", uid).Take(&Student{}).Update(&s)
	if db.Error != nil {
		return &Student{}, db.Error
	}
	//For display Updated Student
	err := db.Debug().Model(&Student{}).Where("user_id = ?", uid).Take(&s).Error
	if err != nil {
		return &Student{}, err
	}
	return s, nil
}
