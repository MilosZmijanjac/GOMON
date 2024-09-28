package models

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     uint8  `json:"role"`
	IsActive bool   `json:"isActive"`
}
type User struct {
	Username string
	Password string
	Role     uint8
	IsActive bool
}
type RegisterResponse struct{}

type ListRequest struct{}

type ListResponse struct {
	Users []User `json:"users"`
}
type UpdateRequest struct {
	Username string `json:"username"`
	Role     uint8  `json:"role"`
	IsActive bool   `json:"isActive"`
}

type UpdateResponse struct{}
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
