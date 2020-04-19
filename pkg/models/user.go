package models

import (
	"html"
	"strings"
	"time"

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

	//hashedPassword, _ := Hash(u.Password)
	//u.Password = hashedPassword

}
