package gorm_client

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name    string
	Email   string
	Address string
	Age     int
}
