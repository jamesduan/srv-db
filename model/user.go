package model

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model

	// ID     int    `gorm:"AUTO_INCREMENT"`
	Name   string `gorm:"size:255"`
	Number string

	Sex string
	Age int

	Emails []Email

	HomeAdresss      []HomeAdress
	ShoppingAdresses []ShoppingAdress
	CreditCards      []CreditCard
	AliPay           AliPay
	Wechat           Wechat
}

type HomeAdress struct {
	gorm.Model
	UsersID int
	Address string         `gorm:"type:varchar(100);unique"`
	Post    sql.NullString `gorm:"not null"`
}

type Email struct {
	gorm.Model
	UsersID    int
	Email      string
	Subscribed bool
}

type ShoppingAdress struct {
	gorm.Model
	UsersID    int
	Province   string
	City       string
	AddrDetail string
	Phone      string
}

type CreditCard struct {
	gorm.Model
	UsersID uint
	Number  string
}

type AliPay struct {
	gorm.Model
	Name     string
	NickName string
	Number   string
}

type Wechat struct {
	gorm.Model
	Name     string
	NickName string
}
