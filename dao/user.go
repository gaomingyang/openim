package dao

import (
	"crypto/md5"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
)

type User struct {
	Id       int64  `db:"id" json:"id"`
	UserName string `db:"user_name" json:"user_name"`
	Password string `db:"password" json:"password"`
}

func UserRegister(userName string, password string) (err error) {
	md5Password := getMd5(password)
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
	md5Password := getMd5(password)
	err = DBInstance.Table("users").Where("user_name = ? and password = ?", userName, md5Password).First(&user).Error
	return
}

func getMd5(inputStr string) string {
	h := md5.New()
	io.WriteString(h, inputStr)
	md5Str := fmt.Sprintf("%x", h.Sum(nil))
	return md5Str
}
