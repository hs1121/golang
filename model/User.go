package model

type Address struct {
	City    string `gorm:"type:varchar(100)"`
	State   string `gorm:"type:varchar(100)"`
	Pincode string `gorm:"type:varchar(20)"`
}

type User struct {
	ID      int   `gorm:"primaryKey"`
	Name    string  `gorm:"type:varchar(100)"`
	Age     int     `gorm:"type:int"`
	Address Address `gorm:"embedded"`
}
