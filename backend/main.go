package main

import (
	"app/config"
	"app/controllers"
	"log"
)

func main() {
	// 初始化数据库
	dsn := config.SetDsn()
	db, err := config.InitDatabase(dsn)
	//user := models.User{
	//	Username: "lps",
	//	Password: "666",
	//}
	//db.Create(&user)  创建成功
	if err != nil {
		log.Fatalf("无法连接数据库: %s\n", err)
	}
	// 初始化路由函数
	r := config.InitGinEngine(db)
	r.POST("/user/register/", controllers.Register)
	r.POST("/user/login", controllers.Login)
	err = r.Run(":8080")
	if err != nil {
		panic("启动服务失败")
	}
}
