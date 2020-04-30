package models

import (
	"time"
)

//Student struct
type Student struct {
	StudentID   uint32 `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	StudentName string `gorm:"size:40;not null;" json:"username"`
	ClassID     string `gorm:"size:50;not null;unique" json:"email"`

	CreatedBy int       `json:"_"`
	CreatedOn time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"_"`
	UpdatedBy int       `json:"_"`
	UpdatedOn time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"_"`
}
