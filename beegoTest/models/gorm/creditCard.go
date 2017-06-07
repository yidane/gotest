package gorm

import (
	"github.com/jinzhu/gorm"
)

//CreditCard for what
type CreditCard struct {
	gorm.Model
	UserID uint
	Number string
}
