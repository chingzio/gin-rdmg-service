package models

import (
	"errors"
	"fmt"
	"log"
)

type User struct {
	Id       uint   `gorm:"column:id;type:int;not null"`
	Username string `gorm:"column:username;type:varchar(128);not null"`
	Password string `gorm:"column:password;type:varchar(32);not null"`
}

type UserRole struct {
	Id       uint   `gorm:"column:id;type:int"`
	UserId   uint   `gorm:"column:user_id;type:int"`
	RoleName string `gorm:"column:role_name;type:varchar(128)"`
}

type UserLogin struct {
	Username string `json:"username" example:"admin"`
	Password string `json:"password" example:"123456"`
}

func SelectOneByName(name string) (User, error) {
	var user User
	Db.Where("username = ?", name).First(&user)
	if (User{} == user) {
		fmt.Println("未查询到此用户" + name)
		return user, errors.New("未查询到此用户")
	}
	return user, nil
}

func SelectUserRoles(name string) ([]string, error) {
	var userRole []UserRole
	err := Db.Raw("SELECT ur.role_name FROM user_role ur, user u WHERE u.username = ? AND ur.user_id = u.id", name).Scan(&userRole).Error
	if err != nil {
		log.Println(err)
	}
	var roleList []string
	for i := 0; i < len(userRole); i++ {
		roleList = append(roleList, userRole[i].RoleName)
	}
	return roleList, err
}

func SelectPwd(name string) string {
	user, err := SelectOneByName(name)
	if err != nil {
		return ""
	} else {
		return user.Password
	}
}
