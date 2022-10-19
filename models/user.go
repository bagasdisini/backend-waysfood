package models

type User struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Password string `json:"password" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" gorm:"type: varchar(255)"`
	Location string `json:"location" gorm:"type: varchar(255)"`
	Image    string `json:"image" gorm:"type: varchar(255)"`
	Role     string `json:"role" gorm:"type: varchar(255)"`
	Gender   string `json:"gender" gorm:"type: varchar(255)"`
}

type UsersProfileResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
