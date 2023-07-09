package models

import "gorm.io/gorm"

type Books struct {
	gorm.Model
	Title string `db:"title" gorm:"not null"`
}
