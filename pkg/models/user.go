package models

import (
	"errors"
	"html"
	"regexp"
	"strings"
	"time"

	"github.com/jinzhu/gorm"

	"golang.org/x/crypto/bcrypt"
)

//User struct
type User struct {
	UserID      int       `gorm:"primary_key;AUTO_INCREMENT" json:"UserID"`
	UserName    string    `gorm:"size:40;not null;" json:"UserName"`
	RoleID      int       `gorm:"not null;" json:"RoleID"`
	PhoneNumber string    `gorm:"size:20;not null;" json:"PhoneNumber"`
	Email       string    `gorm:"size:50;not null;unique" json:"Email"`
	DOB         time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"DOB"`
	Gender      string    `gorm:"size:10;not null" json:"Gender"`
	Password    string    `gorm:"size:100;not null;"`
	IsDeleted   bool      `gorm:"default:false" json:"-"`
	CreatedBy   int       `json:"-"`
	CreatedOn   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedBy   int       `json:"-"`
	UpdatedOn   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
	Roles       Role      `gorm:"foreignkey:RoleID;association_foreignkey:RoleID"`
}

// Message struct
type Message struct {
	Message  string
	Token    string
	UserID   int
	UserName string
}

//Profile Struct
type Profile struct {
	UserID   int
	UserName string
	UserRole int
	Email    string
	Gender   string
}

//Hash function
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//VerifyPassword function
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

//Validate method
func (u *User) Validate(action string) error {

	//Regular expression for email
	emailRe := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	//Regular expression for Date
	//dobRe := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)^(0?[1-9]|1[012]):([0-5][0-9])[ap]m$")

	switch strings.ToLower(action) {
	case "create":
		if u.UserName == "" {
			return errors.New("Required UserName")
		}
		if len(u.UserName) < 3 {
			return errors.New("Username requires minimum 8 characters")

		}
		if u.RoleID == 0 {
			return errors.New("Required RoleID")
		}
		if u.PhoneNumber == "" {
			return errors.New("Required PhoneNumber")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if emailRe.MatchString(u.Email) == false {
			return errors.New("Invalid Email")

		}
		// if dobRe.MatchString(u.DOB.String()) == false {
		// 	return errors.New("Invalid Date Format")

		// }
		if u.Gender == "" {
			return errors.New("Required Gender")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		return nil

	default:
		if u.UserID == 0 {
			return errors.New("Required UserID")
		}
		return nil

	}

}

//ValidateID method
func (u *User) ValidateID(id int) error {

	if id == 0 {
		return errors.New("Required UserID")

	}
	return nil
}

//BeforeSave function
func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

//Prepare function
func (u *User) Prepare() {
	u.UserName = html.EscapeString(strings.TrimSpace(u.UserName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	// hashedPassword, _ := Hash(u.Password)
	// u.Password = string(hashedPassword)
	u.CreatedOn = time.Now()
	u.UpdatedOn = time.Now()
}

//SaveUser method
func (u *User) SaveUser(db *gorm.DB) (*User, error) {

	var err error
	err = db.Debug().Model(&User{}).Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil

}

//FindAllUsers method
func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Preload("Roles").Model(&User{}).Where("is_deleted = ?", false).Order("user_id").Find(&users).Error

	if err != nil {
		return &[]User{}, err
	}
	return &users, nil

}

//FindUserByID method
func (u *User) FindUserByID(db *gorm.DB, uid int) (*User, error) {
	var err error

	err = db.Debug().Model(&User{}).Where("user_id=?", uid).Take(u).Error

	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, nil

}

//UpdateUser a user
func (u *User) UpdateUser(db *gorm.DB, uid int) (*User, error) {
	u.Prepare()
	db = db.Debug().Model(&User{}).Where("user_id = ?", uid).Take(&User{}).Update(&u)
	if db.Error != nil {
		return &User{}, db.Error
	}
	//For display Updated User
	err := db.Debug().Model(&User{}).Where("user_id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

//DeleteUser function
func (u *User) DeleteUser(db *gorm.DB, uid int) (int64, error) {
	db = db.Debug().Model(&User{}).Where("user_id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"is_deleted": true,
		},
	)

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
