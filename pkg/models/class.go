package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

//Class Struct
type Class struct {
	ClassID  uint32    `gorm:"primary_key;AUTO_INCREMENT" json:"ClassID"`
	Standard string    `gorm:"size:50;not null;" json:"Standard"`
	Division string    `gorm:"size:40;not null;" json:"Division"`
	Year     int       `gorm:"not null;" json:"Year"`
	Teacher  []Teacher `gorm:"foreignkey:ClassID;association_foreignkey:ClassID" json:"Teacher"`
}

//SaveClass method
func (c *Class) SaveClass(db *gorm.DB) (*Class, error) {

	var err error
	err = db.Debug().Create(&c).Error
	if err != nil {
		return &Class{}, err
	}
	return c, nil

}

//FindAllClasses method
func (c *Class) FindAllClasses(db *gorm.DB) (*[]Class, error) {
	var err error
	classes := []Class{}
	err = db.Debug().Preload("Teacher").Model(&Class{}).Limit(100).Find(&classes).Error

	if err != nil {
		return &[]Class{}, err
	}
	return &classes, nil

}

//FindClassByID method
func (c *Class) FindClassByID(db *gorm.DB, cid uint32) (*Class, error) {
	var err error

	err = db.Debug().Model(&Class{}).Where("class_id = ?", cid).Take(c).Error

	if err != nil {
		return &Class{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Class{}, errors.New("Class Not Found")
	}
	return c, nil

}

//UpdateClass a user
func (c *Class) UpdateClass(db *gorm.DB, cid uint32) (*Class, error) {
	db = db.Debug().Model(&Class{}).Where("class_id = ?", cid).Take(&Class{}).Update(&c)
	if db.Error != nil {
		return &Class{}, db.Error
	}
	//For display Updated Class
	err := db.Debug().Model(&Class{}).Where("clas_id = ?", cid).Take(&c).Error
	if err != nil {
		return &Class{}, err
	}
	return c, nil
}

//DeleteClass
func (c *Class) DeleteClass(db *gorm.DB, cid uint32) (int64, error) {
	db = db.Debug().Model(&Class{}).Where("class_id = ?", cid).Take(&Class{}).UpdateColumns(
		map[string]interface{}{
			"is_deleted": true,
		},
	)

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
