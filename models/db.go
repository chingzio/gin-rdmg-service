package models

import (
	"fmt"
	"gin-rdmg-service/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

// 初始化
func init() {
	var err error
	constr := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", utils.DbUser, utils.DbPassWord, utils.DbHost, utils.DbPort, utils.DbName)
	fmt.Println("open db...")
	Db, err = gorm.Open("mysql", constr)
	// 不使用复数表名
	Db.SingularTable(true)
	if err != nil {
		fmt.Println(err)
		panic("数据库连接失败")
	}
	// 初始化
	Db.AutoMigrate(&User{}, &UserRole{})
	var count int
	Db.Model(&User{}).Where("id = ?", 1).Or("username = ?", "admin").Count(&count)
	if count == 0 {
		Db.Save(User{1, "admin", "e10adc3949ba59abbe56e057f20f883e"})
	}
	Db.Model(&UserRole{}).Where("id = ?", 1).Or("user_id = ?", 1).Count(&count)
	if count == 0 {
		Db.Save(UserRole{1, 1, "admin"})
	}
	fmt.Println("open db success")
}

func CloseDb() {
	fmt.Println("close db...")
	Db.Close()
	fmt.Println("close db done")
}
