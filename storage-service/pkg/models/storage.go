package models

import "gorm.io/gorm"

type Storage struct{
	gorm.Model
	Name string	`gorm:"column:name;unique;not null"`
	UserID uint64 `gorm:"column:user_id;not null;index"`

	Cells []Cell `gorm:"foreignKey:StorageID"`
}