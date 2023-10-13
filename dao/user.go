package dao

import (
	"errors"
	"openim/common"
	"openim/common/define"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int64  `db:"id" json:"id"`
	UserName string `db:"user_name" json:"user_name"`
	Password string `db:"password" json:"password"`
	Status   int    `db:"status" json:"status"`
}

func UserRegister(userName string, password string) (err error) {
	md5Password := common.GetMD5(password)
	user := User{UserName: userName, Password: md5Password}
	result := DBInstance.Create(&user)
	err = result.Error
	return
}

func UserFindByName(userName string) (num int64, err error) {
	res := DBInstance.Table("users").Where("user_name = ?", userName).Count(&num)
	err = res.Error
	return
}

func UserLogin(userName string, password string) (user User, err error) {
	md5Password := common.GetMD5(password)
	err = DBInstance.Table("users").Select("id", "user_name", "password", "status").Where("user_name = ?", userName).First(&user).Error
	if err != nil {
		return
	}

	if user.Status != define.USER_STATUS_NORMAL {
		err = errors.New("用户状态异常")
		return
	}

	if user.Password != md5Password {
		err = errors.New("wrong password")
		return
	}
	return
}
