package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Account string
	Password string
	UserId int64
	NickName string
	HeadIndex int32
	RegisterTime primitive.DateTime
	RegisterIp string
	LastLoginTime primitive.DateTime
	LastLoginIp string
	Gender int32
	Score int64
	Status int32
	OnlineStatus int32
}