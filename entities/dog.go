package entities

import "gorm.io/gorm"

type Dog struct {
	gorm.Model
	Name string
	Age  int
}
