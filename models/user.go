package models

type User struct {
	tableName struct{} `pg:"users,select:users,alias:users"`

	Id int64 `pg:"id,pk,notnull"`
	Username string `pg:"username,unique,notnull"`
	HashedPassword string `pg:"hashed_password,notnull"`
	GroupID string `pg:"group_id,notnull"`
	IsAdmin bool `pg:"is_admin,use_zero"`
}