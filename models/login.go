package models

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	HashedPassword string `json:"hashed_password" binding:"required"`
}

type LoginResp struct {
	Username string `json:"username"`
	GroupId string `json:"group_id"`
	IsAdmin bool `json:"is_admin"`
}