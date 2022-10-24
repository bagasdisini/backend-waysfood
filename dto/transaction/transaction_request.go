package transactiondto

type CreateTransactionRequest struct {
	Status    string `json:"status" form:"status" gorm:"type: varchar(255)" validate:"required"`
	Qty       int    `json:"qty" form:"qty" gorm:"type: int" validate:"required"`
	ProductID int    `json:"product_id" form:"product_id" gorm:"type: int" validate:"required"`
}

type UpdateTransactionRequest struct {
	Status    string `json:"status" form:"status" gorm:"type: varchar(255)"`
	Qty       int    `json:"qty" form:"qty" gorm:"type: int"`
	ProductID int    `json:"product_id" form:"product_id" gorm:"type: int"`
}
