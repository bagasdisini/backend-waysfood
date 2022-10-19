package usersdto

type CreateUserRequest struct {
	FullName string `json:"fullName" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Gender   string `json:"gender" form:"gender" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Role     string `json:"role" form:"role" validate:"required"`
}

type UpdateUserRequest struct {
	FullName string `json:"fullName" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Gender   string `json:"gender" form:"gender"`
	Phone    string `json:"phone" form:"phone"`
	Role     string `json:"role" form:"role"`
}
