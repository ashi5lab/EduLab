package models

import (
	"errors"
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"

	"golang.org/x/crypto/bcrypt"
)

//User struct
type User struct {
	UserID    uint32    `gorm:"primary_key;auto_increment" json:"id"`
	UserName  string    `gorm:"size:40;not null;unique" json:"username"`
	Email     string    `gorm:"size:50;not null;unique" json:"email"`
	RoleID    int       `gorm:"not null;" json:"role"`
	Password  string    `gorm:"size:100;not null;" json:"_"`
	IsDeleted bool      `gorm:"default:false" json:"_"`
	CreatedBy int       `json:"_"`
	CreatedOn time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"_"`
	UpdatedBy int       `json:"_"`
	UpdatedOn time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"_"`
}

//Hash function
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
func VerifyPassword(hashedPassword, password string) error {
	fmt.Println(Hash(password))
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

//Prepare function
func (u *User) Prepare() {
	u.UserName = html.EscapeString(strings.TrimSpace(u.UserName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	hashedPassword, _ := Hash(u.Password)
	u.Password = string(hashedPassword)

}

//SaveUser method
func (u *User) SaveUser(db *gorm.DB) (*User, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil

}

//FindAllUsers method
func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Where("is_deleted = ?", false).Find(&users).Error

	if err != nil {
		return &[]User{}, err
	}
	return &users, nil

}

//FindUserByID method
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

//UpdateUser a user
func (u *User) UpdateUser(db *gorm.DB, uid uint32) (*User, error) {
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

//DeleteUser
func (u *User) DeleteUser(db *gorm.DB, uid uint32) (int64, error) {
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
