package controllers

import (
	"app/models"
	"app/services"
	"app/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register 处理用户注册的API请求
func Register(c *gin.Context) {
	var user models.User
	user.Username = c.Query("username")
	user.Password = c.Query("password")
	if len(user.Username) < 6 || len(user.Password) < 6 || len(user.Username) > 25 || len(user.Password) > 25 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": 1,
			"status_msg":  "用户名和密码不得为空，且长度应在 6 - 25 个字符之间。",
			"user_id":     nil,
			"username":    user.Username,
		})
		fmt.Println(http.StatusBadRequest, "用户名和密码不得为空，且长度应在 6 - 25 个字符之间。")
		return
	}
	// 调用service层的注册方法
	err := services.UserRegister(&user, c)
	if err != nil {
		fmt.Println(user.ID, user.Username)
		return
	} else {
		// 生成新Token
		newToken, err := utils.GenerateToken(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status_code": 1,
				"status_msg":  "生成token令牌失败。",
				"user_id":     user.ID,
			})
			fmt.Println(http.StatusInternalServerError, "生成token令牌失败。")
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status_code": 0,
			"status_msg":  "注册成功!",
			"user_id":     user.ID,
			"token":       newToken,
		})
		fmt.Println(http.StatusCreated, "注册成功!")
		fmt.Println(user.ID, user.Username, newToken)
	}

}

// Login 处理用户登录请求
func Login(c *gin.Context) {
	var user models.User
	var inputUser models.User

	// 提取用户名和密码
	inputUser.Username = c.Query("username")
	inputUser.Password = c.Query("password")

	// 验证用户名密码非空
	if inputUser.Username == "" || inputUser.Password == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status_code": 1,
			"status_msg":  "没填写用户名或者密码",
		})
		fmt.Println(http.StatusUnauthorized, "没填写用户名或者密码")
		return
	}

	err := services.UserLoginVerify(&user, &inputUser, c)
	if err != nil {
		return
	}

	// 生成新Token
	newToken, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": 1,
			"status_msg":  "生成token令牌失败。",
			"user_id":     user.ID,
		})
		fmt.Println(http.StatusInternalServerError, "生成token令牌失败。")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "登录成功",
		"user_id":     user.ID,
		"token":       newToken,
	})
	fmt.Println(http.StatusOK, "登录成功")
}
