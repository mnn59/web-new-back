package models

type User struct {
	Id uint
	Name string
	FamilyName string
	Email string `gorm:"unique"`
	Password []byte
	Address string
	ChargeAmount string
}
