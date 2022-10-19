package authdto

type AuthRequest struct {
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Password string `gorm:"type: varchar(255)" json:"password"`
}
