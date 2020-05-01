package models

type Role struct {
	RoleID   uint32 `gorm:"primary_key;" json:"roleid"`
	RoleName string `gorm:"size:40;not null;" json:"rolename"`
}
