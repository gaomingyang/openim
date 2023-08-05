package services

import (
	"gorm.io/gorm"
	"openim/dao"
)

// 注入一个db
type FriendService struct {
	DB *gorm.DB
}

func NewFriendService() *FriendService {
	return &FriendService{
		DB: dao.DBInstance,
	}
}

func (s FriendService) list() {

}
