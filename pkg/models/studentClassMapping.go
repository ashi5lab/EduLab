package models

import (
	"time"
)

//StudentClassMapping struct
type StudentClassMapping struct {
	StudentClassID int       `gorm:"primary_key;AUTO_INCREMENT" json:"StudentClassID"`
	StudentID      int       `gorm:"StudentID:UserID;association_foreignkey:StudentID" json:"StudentID"`
	ClassID        int       `gorm:"ClassID:UserID;association_foreignkey:ClassID" json:"ClassID"`
	CreatedBy      int       `json:"_"`
	CreatedOn      time.Time `gorm:"default:CURRENT_TIMESTAMP" `
	UpdatedBy      int
	UpdatedOn      time.Time `gorm:"default:CURRENT_TIMESTAMP" `
}
