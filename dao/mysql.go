package dao

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 连接数据库
func InitMySQL() (err error) {
	dsn := "root:028135Cc@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	p, _ := DB.DB()
	return p.Ping()
}

func Close() {
	//获取通用数据库对象 sql.DB，然后使用其提供的功能
	sqlDB, err := DB.DB()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	sqlDB.Close() //数据库关闭
}
