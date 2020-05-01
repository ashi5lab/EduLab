package models

type Class struct {
	ClassID   uint32 `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Standard  string `gorm:"size:50;not null;" json:"standard"`
	Division  string `gorm:"size:40;not null;" json:"Division"`
	Year      int    `gorm:"not null;" json:"year"`
	TeacherID int    `gorm:"foreignkey:TeacherID;association_foreignkey:TeacherID" json:"TeacherID"`
}
