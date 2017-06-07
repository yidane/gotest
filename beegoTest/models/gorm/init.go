package gorm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func open() (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", "root:sasasasasa@tcp(localhost:3306)/test?charset=utf8")
	return db, nil
}

func Init() error {
	db, err := open()
	if err != nil {
		return err
	}
	defer db.Close()

	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	// 	return "T_" + defaultTableName
	// }

	db.AutoMigrate(&CreditCard{}, &Email{}, &Language{}, &Address{}, &User{})
	return nil
}

func HasUerTable() bool {
	db, err := open()
	if err != nil {
		return false
	}
	defer db.Close()
	flag := db.HasTable(&Address{})

	return flag
}
