package models

//Role struct
type Role struct {
	RoleID   uint32 `gorm:"primary_key;" json:"RoleID"`
	RoleName string `gorm:"size:20;not null;" json:"RoleName"`
}

//ROLE values
/*
1 -> Admin
2 -> Student
3 -> Teacher

*/
