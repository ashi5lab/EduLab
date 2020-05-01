package models

//Teacher struct
type Teacher struct {
	TeacherID     uint32 `gorm:"primary_key;AUTO_INCREMENT" json:"TeacherID"`
	Qualification string `gorm:"size:30;" json:"Qualification"`
	Subject       string `gorm:"size:20;" json:"Subject"`
}
