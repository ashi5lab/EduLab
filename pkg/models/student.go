package models

import (
	"time"
)

//Student struct
type Student struct {
	StudentID                 int       `gorm:"primary_key;AUTO_INCREMENT" json:"StudentID"`
	UserID                    int       `gorm:"foreignkey:UserID;association_foreignkey:UserID" json:"UserID"`
	StudentAdmno              int       `gorm:"not null;unique" json:"StudentAdmno"`
	StudentSlno               int       `gorm:"not null;unique" json:"StudentSlno"`
	StudentAppno              int       `gorm:"not null;unique" json:"StudentAppno"`
	StudentGuardianName       string    `gorm:"size:50;" json:"StudentGuardianName"`
	StudentGuardianOccupation string    `gorm:"size:50;" json:"StudentGuardianOccupation"`
	StudentGuardianRelation   string    `gorm:"size:50;" json:"StudentGuardianRelation"`
	StudentAddress            string    `gorm:"size:100;" json:"StudentAddress"`
	StudentReligion           string    `gorm:"size:20;" json:"StudentReligion"`
	StudentCaste              string    `gorm:"size:20;" json:"StudentCaste"`
	StudentCategory           string    `gorm:"size:20;" json:"StudentCategory"`
	StudentIsOEC              bool      `json:"StudentIsOEC"`
	StudentIsLingMin          bool      `json:"StudentIsLingMin"`
	StudentLingMinDesc        string    `gorm:"size:50;" json:"StudentLingMinDesc"`
	StudentDOA                time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"StudentDOA"`
	StudentAdmissionCategory  string    `gorm:"size:50;" json:"StudentAdmissionCategory"`
	CreatedBy                 int       `json:"_"`
	CreatedOn                 time.Time `gorm:"default:CURRENT_TIMESTAMP" json:""`
	UpdatedBy                 int       `json:""`
	UpdatedOn                 time.Time `gorm:"default:CURRENT_TIMESTAMP" json:""`
}
