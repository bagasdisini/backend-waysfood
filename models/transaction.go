package models

type Transaction struct {
	ID int `json:"id" gorm:"primary_key:auto_increment"`
	// Users     UsersProfileResponse `json:"users" gorm:"foreignKey:UsersProfileResponse"`
	Status    string          `json:"status"`
	ProductID int             `json:"product_id"`
	Product   ProductResponse `json:"product"`
}
