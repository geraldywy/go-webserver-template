package models

type SignupReq struct {
	Username string `json:"username" binding:"required"`
	HashedPassword string `json:"hashed_password" binding:"required"`
	GroupID string `json:"group_id" binding:"required"`
	IsAdmin *bool `json:"is_admin" binding:"required"`
}
