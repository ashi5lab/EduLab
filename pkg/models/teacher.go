package models

//Teacher struct
type Teacher struct {
	TeacherID     uint32 `gorm:"primary_key;AUTO_INCREMENT" json:"TeacherID"`
	UserID        int    `gorm:"foreignkey:UserID;association_foreignkey:UserID" json:"UserID"`
	Qualification string `gorm:"size:30;" json:"Qualification"`
	Subject       string `gorm:"size:20;" json:"Subject"`
	ClassID       int    `gorm:""`
}
