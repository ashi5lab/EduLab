package models

type Class struct {
	ClassID   uint32 `gorm:"primary_key;AUTO_INCREMENT" json:"ClassID"`
	Standard  string `gorm:"size:50;not null;" json:"Standard"`
	Division  string `gorm:"size:40;not null;" json:"Division"`
	Year      int    `gorm:"not null;" json:"Year"`
	TeacherID int    `gorm:"foreignkey:TeacherID;association_foreignkey:TeacherID" json:"TeacherID"`
}
