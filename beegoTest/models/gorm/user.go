package gorm

import (
	"time"

	"database/sql"

	"github.com/jinzhu/gorm"
)

//User define for description user information
type User struct {
	gorm.Model
	Birthday          time.Time
	Age               int
	Name              string `gorm:"size:255"`
	Num               int    `gorm:"AUTO_INCREMENT"`
	CreditCard        CreditCard
	Emails            []Email
	BillingAddress    Address
	BillingAddressID  sql.NullInt64
	ShippingAddress   Address
	ShippingAddressID sql.NullInt64
	IgnoreMe          int `gorm:"-"`
	Language          []Language
}

//TableName set the table name for object User
func (User) TableName() string {
	return "t_Users"
}
