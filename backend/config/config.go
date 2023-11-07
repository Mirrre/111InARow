package config

import (
	"app/models"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/redis/go-redis/v9"
)

// InitDatabase 通过dsn来初始化db链接
func InitDatabase(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// 自动将表单模型结构体迁移成数据库表单
	err = db.AutoMigrate(&models.User{}, &models.Relation{})
	if err != nil {
		return nil, err
	}

	err = db.Table("videos").AutoMigrate(&models.Video{})
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

	// 添加CORS中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},                                       // 允许所有来源，您可以根据需要指定允许的来源
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},            // 允许的HTTP方法
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"}, // 允许的标头
	}))

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	return r
}

var Rdb *redis.Client

// InitRedis 初始化redis
func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "r-7xvwyunrh3rymq6515pd.redis.rds.aliyuncs.com:6379",
		Username: "lips",
		Password: "yhy@PS4611",
		DB:       0, // 默认DB 0
	})
}
