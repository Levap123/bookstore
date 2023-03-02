package dto

type SignUpDTO struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type SignInDTO struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type UpdateUserDTO struct {
	Email       string `json:"email,omitempty"`
	Username    string `json:"username,omitempty"`
	NewPassword string `json:"new_password,omitempty"`
	OldPassword string `json:"old_password,omitempty"`
}
