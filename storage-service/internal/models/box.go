package models

import "gorm.io/gorm"

type Box struct{
	gorm.Model
	Name string `gorm:"column:name;unique;not null"`
}