package config

import (
	"fmt"
	"github.com/spf13/viper"
	"goblog/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB

// LoadConfig 加载配置文件

func InitMySQL() {
	// 读取配置文件中的数据库信息
	LoadConfig()

	mysqlConfig := viper.Sub("mysql")
	if mysqlConfig == nil {
		panic("MySQL 配置不存在")
	}

	user := mysqlConfig.GetString("user")
	password := mysqlConfig.GetString("password")
	database := mysqlConfig.GetString("database")
	host := mysqlConfig.GetString("host")
	port := mysqlConfig.GetString("port")
	charset := mysqlConfig.GetString("charset")

	// 格式化数据库连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		user, password, host, port, database, charset)
	fmt.Println(dsn)

	var err error
	MysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database!")
	}

	MysqlDB.AutoMigrate(&models.Post{})
	MysqlDB.AutoMigrate(&models.User{})
}
