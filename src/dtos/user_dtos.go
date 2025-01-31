package dtos

type CreateUserDto struct {
	Username string `json:"username" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

var CreateUserDtoValidationMessages = map[string]string{
	"Username": "Username is required and must be at least 3 characters",
	"Email":    "Email is required and must be a valid email address",
	"Password": "Password is required and must be at least 6 characters",
}
