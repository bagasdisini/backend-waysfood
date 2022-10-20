package usersdto

type UserResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Location string `json:"location"`
	Image    string `json:"image"`
	Role     string `json:"role"`
}
