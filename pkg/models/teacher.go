package models

type Teacher struct {
	TeacherID     uint32 `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Qualification string `gorm:"size:30;" json:"qualification"`
	Subject       string `gorm:"size:40;" json:"subject"`
}
