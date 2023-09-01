package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Address struct {
	gorm.Model
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
}

type Contact struct {
	gorm.Model
	Phone     string   `json:"phone"`
	Email     string   `json:"email"`
	Address   *Address `json:"address" gorm:"foreignkey:AddressID"`
	AddressID uint     `json:"-"`
}

type Employee struct {
	gorm.Model
	Name      string   `json:"name"`
	Position  string   `json:"position"`
	Contact   *Contact `json:"contact" gorm:"foreignkey:ContactID"`
	ContactID uint     `json:"-"`
}

