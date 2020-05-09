package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

//Student struct
type Student struct {
	StudentID                 int       `gorm:"primary_key;AUTO_INCREMENT" json:"StudentID"`
	UserID                    int       `gorm:"not null;"`
	StudentAdmno              int       `gorm:"not null;unique" json:"StudentAdmno"`
	StudentSlno               int       `gorm:"not null;unique" json:"StudentSlno"`
	StudentAppno              int       `gorm:"not null;unique" json:"StudentAppno"`
	StudentGuardianName       string    `gorm:"size:50;" json:"StudentGuardianName"`
	StudentGuardianOccupation string    `gorm:"size:50;" json:"StudentGuardianOccupation"`
	StudentGuardianRelation   string    `gorm:"size:50;" json:"StudentGuardianRelation"`
	StudentAddress            string    `gorm:"size:100;" json:"StudentAddress"`
	StudentReligion           string    `gorm:"size:20;" json:"StudentReligion"`
	StudentCaste              string    `gorm:"size:20;" json:"StudentCaste"`
	StudentCategory           string    `gorm:"size:20;" json:"StudentCategory"`
	StudentIsOEC              bool      `json:"StudentIsOEC"`
	StudentIsLingMin          bool      `json:"StudentIsLingMin"`
	StudentLingMinDesc        string    `gorm:"size:50;" json:"StudentLingMinDesc"`
	StudentDOA                time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"StudentDOA"`
	StudentAdmissionCategory  string    `gorm:"size:50;" json:"StudentAdmissionCategory"`
	IsDeleted                 bool      `gorm:"default:false" json:"-"`
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
	err = db.Debug().Preload("Users", "is_deleted=?", false).Model(&Student{}).Limit(100).Find(&students).Error

	if err != nil {
		return &[]Student{}, err
	}
	return &students, nil

}

//FindStudentByID method
func (s *Student) FindStudentByID(db *gorm.DB, sid uint32) (*Student, error) {
	var err error

	err = db.Debug().Model(&Student{}).Where("student_id=?", sid).Take(s).Error

	if err != nil {
		return &Student{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Student{}, errors.New("Student Not Found")
	}
	return s, nil

}

//UpdateStudent a student
func (s *Student) UpdateStudent(db *gorm.DB, sid uint32) (*Student, error) {
	s.Prepare()
	db = db.Debug().Model(&Student{}).Where("student_id = ?", sid).Take(&Student{}).Update(&s)
	if db.Error != nil {
		return &Student{}, db.Error
	}
	//For display Updated Student
	err := db.Debug().Model(&Student{}).Where("student_id = ?", sid).Take(&s).Error
	if err != nil {
		return &Student{}, err
	}
	return s, nil
}
