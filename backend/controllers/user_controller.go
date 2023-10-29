package controllers

import (
	"app/models"
	"app/responses"
	"app/services"
	"app/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register 处理用户注册的API请求
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, responses.RegisterResponse{
			StatusCode: 1,
			StatusMsg:  "接收用户注册数据失败",
			UserID:     0,
			Token:      "",
		})
	}
	if len(user.Username) < 6 || len(user.Password) < 6 || len(user.Username) > 25 || len(user.Password) > 25 {
		c.JSON(http.StatusBadRequest, responses.RegisterResponse{
			StatusCode: 1,
			StatusMsg:  "用户名和密码不得为空，且长度应在 6 - 25 个字符之间。",
			UserID:     0,
			Token:      "",
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
			c.JSON(http.StatusInternalServerError, responses.RegisterResponse{
				StatusCode: 1,
				StatusMsg:  "生成token令牌失败。",
				UserID:     0,
				Token:      "",
			})
			fmt.Println(http.StatusInternalServerError, "生成token令牌失败。")
			return
		}

		c.JSON(http.StatusCreated, responses.RegisterResponse{
			StatusCode: 0,
			StatusMsg:  "注册成功",
			UserID:     user.ID,
			Token:      newToken,
		})
		fmt.Println(http.StatusCreated, "注册成功!")
		fmt.Println(user.ID, user.Username, newToken)
	}

}

// Login 处理用户登录请求
func Login(c *gin.Context) {
	var user models.User
	var inputUser models.User
	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(http.StatusBadRequest, responses.RegisterResponse{
			StatusCode: 1,
			StatusMsg:  "接收用户注册数据失败",
			UserID:     0,
			Token:      "",
		})
	}
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
		c.JSON(http.StatusInternalServerError, responses.LoginResponse{
			StatusCode: 1,
			StatusMsg:  "生成token令牌失败。",
			Token:      "",
			UserID:     user.ID,
		})
		fmt.Println(http.StatusInternalServerError, "生成token令牌失败。")
		return
	}

	c.JSON(http.StatusOK, responses.LoginResponse{
		StatusCode: 0,
		StatusMsg:  "登录成功",
		Token:      newToken,
		UserID:     user.ID,
	})
	fmt.Println(http.StatusOK, "登录成功")
}

type UserId struct {
	UserId int `json:"user_id"`
}

func GetUser(c *gin.Context) {
	var userId UserId
	err := c.ShouldBindJSON(userId) //从json中获取user_id

	if err != nil || userId.UserId <= 0 {
		c.JSON(http.StatusBadRequest, responses.UserResponse{
			StatusCode: 1,
			StatusMsg:  "目标用户 ID 无效。",
			User:       responses.User{},
		})
		return
	}
	var user models.User
	// 获取用户信息
	isFollowed, err := services.GetUser(&user, userId.UserId, c)
	if err == nil {
		userResponse := responses.User{
			ID:              user.ID,
			Name:            user.Username,
			FollowCount:     user.Profile.FollowCount,
			IsFollowed:      isFollowed,
			Avatar:          user.Profile.Avatar,
			BackgroundImage: user.Profile.Background,
			Signature:       user.Profile.Signature,
			TotalFavorited:  user.Profile.TotalFavorited,
			WorkCount:       user.Profile.WorkCount,
			FavoriteCount:   user.Profile.FavoriteCount,
		}
		fmt.Println(http.StatusOK, userResponse)
		c.JSON(http.StatusOK, responses.UserResponse{
			StatusCode: 0,
			StatusMsg:  "OK",
			User:       userResponse,
		})
	}

}
