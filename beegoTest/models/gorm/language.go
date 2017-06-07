package gorm

type Language struct {
	ID   int
	Name string `gorm:"index:id_name_code"`
	Code string `gorm:"index:id_name_code"`
}
