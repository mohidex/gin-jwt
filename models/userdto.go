package models

type RequestUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Admin  bool   `json:"is_admin"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
