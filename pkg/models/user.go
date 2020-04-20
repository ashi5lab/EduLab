package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserId    uint32    `gorm:"primary_key;auto_increment" json:"id"`
	UserName  string    `gorm:"size:40;not null;unique" json:"name"`
	Email     string    `gorm:"size:50;not null;unique" json:"email"`
	RoleId    int       `gorm:"not null;unique" json:"role"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	IsDeleted bool      `gorm:"default:false" gorm:"json:"deleted"`
	CreatedBy int       `gorm:"json:"created_by"`
	CreatedOn time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_on"`
	UpdatedBy int       `gorm:"json:"updated_by"`
	UpdatedOn time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_on"`
}

//Hashing function

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//prepare function
func (u *User) Prepare() {
	u.UserName = html.EscapeString(strings.TrimSpace(u.UserName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	hashedPassword, _ := Hash(u.Password)
	u.Password = string(hashedPassword)

}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil

}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error

	if err != nil {
		return &[]User{}, err
	}
	return &users, nil

}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
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
