package models

type Transaction struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	Qty       int                  `json:"qty" form:"qty" gorm:"type: int"`
	UsersID   int                  `json:"user_id"`
	Users     UsersProfileResponse `json:"users"`
	Status    string               `json:"status" form:"status" gorm:"type: varchar(255)"`
	ProductID int                  `json:"products_id" form:"products_id" gorm:"type: int"`
	Product   ProductResponse      `json:"product"`
}
