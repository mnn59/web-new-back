package models

type User struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	FamilyName   string `json:"family_name"`
	Email        string `json:"email" gorm:"unique"`
	Password     []byte `json:"-"`
	Address      string `json:"address"`
	ChargeAmount string `json:"charge_amount"`
}
