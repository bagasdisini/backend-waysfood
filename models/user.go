package models

type User struct {
	ID       int    `json:"id" gorm:"primary_key:auto_increment;"`
	Email    string `json:"email" form:"email" gorm:"type: varchar(255)"`
	Password string `json:"password" form:"password" gorm:"type: varchar(255)"`
	FullName string `json:"fullName" form:"fullName" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" form:"phone" gorm:"type: varchar(255)"`
	Location string `json:"location" form:"location" gorm:"type: varchar(255)"`
	Image    string `json:"image" form:"image" gorm:"type: varchar(255)"`
	Role     string `json:"role" form:"role" gorm:"type: varchar(255)"`
	Gender   string `json:"gender" form:"gender" gorm:"type: varchar(255)"`
}

type UsersProfileResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName"`
	Location string `json:"location"`
	Email    string `json:"email"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
