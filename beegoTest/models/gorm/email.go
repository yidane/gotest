package gorm

//Email define for email information
type Email struct {
	ID         int
	UserID     int
	Email      string `gorm:"index"`
	Subscribed bool   `gorm:"type:varchar(100);unique_index"`
}
