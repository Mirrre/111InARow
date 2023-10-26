package services

import (
	"app/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

func UserRegister(user *models.User, c *gin.Context) (err error) {
	// 对密码加密处理
	user.Password, err = HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": 1,
			"status_msg":  "密码加密失败",
			"user_id":     nil,
		})
		fmt.Println(http.StatusBadRequest, "密码加密失败:", err)
	}
	// 使用GORM将用户数据存储到数据库中
	db := c.MustGet("db").(*gorm.DB)
	if err = db.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": 1,
			"status_msg":  "注册失败",
			"user_id":     nil,
		})
		fmt.Println(http.StatusBadRequest, "注册失败:", err)
		return err
	}
	return nil
}

// HashPassword 使用 Bcrypt 算法生成密码哈希值
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	//fmt.Println(string(hashedPassword))
	return string(hashedPassword), nil
}

func UserLoginVerify(user *models.User, inputUser *models.User, c *gin.Context) (err error) {
	// 使用GORM检索用户是否存在
	db := c.MustGet("db").(*gorm.DB)
	if err = db.Where("username = ?", inputUser.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status_code": 1,
			"status_msg":  "用户不存在",
		})
		fmt.Println(http.StatusUnauthorized, "用户不存在")
		return err
	}

	// 验证密码 (这里验证的哈希后的密码)
	if ComparePasswords(user.Password, inputUser.Password) {
		return nil
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status_code": 1,
			"status_msg":  "密码错误",
		})
		fmt.Println(http.StatusUnauthorized, "密码错误")
		err = errors.New("密码错误")
		return err
	}
}

// ComparePasswords 比较输入的密码与哈希值是否匹配
func ComparePasswords(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
