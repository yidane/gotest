package gorm

import (
	"database/sql"
)

//Address define for address information
type Address struct {
	ID       int
	Address1 string         `gorm:"not null;unique"`
	Address2 string         `gorm:"type:nvarchar(100);unique"`
	Post     sql.NullString `gorm:"not null"`
}
