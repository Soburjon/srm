package models

type RegisterRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type RegisterResponse struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}

type LoginResponse struct {
	UserID    string `json:"user_id"`
	Authorize string `json:"authorize"`
}

type StandardErrorModel struct {
	Error string `json:"error"`
}

type SuccessMessage struct {
	Success bool `json:"success"`
}
