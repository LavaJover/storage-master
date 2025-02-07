package models

import "gorm.io/gorm"

type Cell struct {
	gorm.Model
	Name string `gorm:"column:name;unique;not null"`
	StorageID uint `gorm:"column:storage_id;not null;index"`

	Boxes []Box `gorm:"foreignKey:CellID"`
}