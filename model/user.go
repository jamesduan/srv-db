package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model
	Name   string `gorm:"size:255"`
	Number string

	Sex      string
	Age      int
	BirthDay time.Time

	Emails []Email

	HomeAddresses    []HomeAddress
	ShoppingAdresses []ShoppingAdress
	CreditCards      []CreditCard
	AliPay           AliPay
	Wechat           Wechat
}

func (Users) TableName() string {
	return "user"
}

type HomeAddress struct {
	gorm.Model
	// ID      int `gorm:"AUTO_INCREMENT"`
	UsersID int
	Address string `gorm:"type:varchar(100);unique"`
	Post    string
}

func (HomeAddress) TableName() string {
	return "home_address"
}

type Email struct {
	gorm.Model
	// ID         int `gorm:"AUTO_INCREMENT"`
	UsersID    int
	Email      string
	Subscribed bool
}

func (Email) TableName() string {
	return "email"
}

type ShoppingAdress struct {
	gorm.Model
	// ID         int `gorm:"AUTO_INCREMENT"`
	UsersID    int
	Province   string
	City       string
	AddrDetail string
	Phone      string
}

func (ShoppingAdress) TableName() string {
	return "shopping_adress"
}

type CreditCard struct {
	gorm.Model
	// ID      int `gorm:"AUTO_INCREMENT"`
	UsersID uint
	Number  string
}

func (CreditCard) TableName() string {
	return "credit_card"
}

type AliPay struct {
	gorm.Model
	// ID       int `gorm:"AUTO_INCREMENT"`
	Name     string
	NickName string
	Number   string
}

func (AliPay) TableName() string {
	return "alipay"
}

type Wechat struct {
	gorm.Model
	// ID       int `gorm:"AUTO_INCREMENT"`

	Name     string
	NickName string
}

func (Wechat) TableName() string {
	return "wechat"
}
