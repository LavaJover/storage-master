package models

import "gorm.io/gorm"

type Box struct{
	gorm.Model
	Name string `gorm:"column:name;unique;not null"`
	CellID uint `gorm:"column:cell_id"`
}