package config

import (
	"app/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDatabase 通过dsn来初始化db链接
func InitDatabase(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// 自动将表单模型结构体迁移成数据库表单
	err = db.AutoMigrate(&models.User{}, &models.UserProfile{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// SetDsn 自动判断测试/生产环境来生成不同的dsn，对应不同的数据库
func SetDsn() string {
	User := "lips"
	Pass := "yhy@PS4611"
	Host := "rm-cn-9lb3g9sil00042do.rwlb.rds.aliyuncs.com"
	Port := "2396"

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/111inarow?charset=utf8mb4&parseTime=True&loc=Local", User, Pass, Host, Port)
}

// InitGinEngine 初始化路由函数
func InitGinEngine(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	return r
}
