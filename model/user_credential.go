package model

type UserCredential struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	Token string `json:"token"`
	UserRole string `json:"role"`
}
