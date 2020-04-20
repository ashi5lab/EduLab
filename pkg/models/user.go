package models

import (
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID    uint32    `gorm:"primary_key;auto_increment" json:"id"`
	UserName  string    `gorm:"size:40;not null;unique" json:"username"`
	Email     string    `gorm:"size:50;not null;unique" json:"email"`
	RoleID    int       `gorm:"not null;unique" json:"role"`
	Password  string    `gorm:"size:100;not null;" json:"_"`
	IsDeleted bool      `gorm:"default:false" json:"_"`
	CreatedBy int       `gorm:json:"_"`
	CreatedOn time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"_"`
	UpdatedBy int       `gorm:json:"_"`
	UpdatedOn time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"_"`
}

//Hashing function

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//prepare function
func (u *User) Prepare() {
	u.UserName = html.EscapeString(strings.TrimSpace(u.UserName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	//hashedPassword, _ := Hash(u.Password)
	//u.Password = hashedPassword

}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil

}
