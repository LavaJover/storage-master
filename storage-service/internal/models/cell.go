package models

import "gorm.io/gorm"

type Cell struct {
	gorm.Model
	Name string `gorm:"column:name;unique;not null"`
}