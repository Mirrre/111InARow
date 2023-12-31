package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtKey = []byte("jwt_secret_key") //

// GenerateToken 生成新Token
func GenerateToken(userID uint) (string, error) {
	// 自定义Token的声明，声明可以理解为一个JSON数据包，包含了我们想要封装在Token里面的信息
	claims := jwt.MapClaims{
		"user_id": userID,
		// exp - 过期时间，格式为Unix时间戳.
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	// 利用claims生成一个Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用秘钥来对Token进行签名
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidateToken 解析和验证Token
func ValidateToken(signedToken string) (uint, error) {
	// 定义一个空的MapClaims，用来保存我们Token中的声明（claims）
	claims := &jwt.MapClaims{}

	// 解析Token，同时将解析出来的claims填入上面声明的空claims，
	// 注意第三个参数是一个函数参数，用来指定解析Token时用什么秘钥
	// 也就是之前定义的jwtKey
	token, err := jwt.ParseWithClaims(
		signedToken,
		claims,
		func(token *jwt.Token) (interface{}, error) { return jwtKey, nil })

	if err != nil || !token.Valid {
		return 0, errors.New("令牌无效")
	}

	// 从解析出来的claims里面提取用户名，并且断言它是字符串类型
	userIDFloat, ok := (*claims)["user_id"].(float64)
	if !ok {
		return 0, errors.New("令牌不包含用户id")
	}

	userID := uint(userIDFloat)
	return userID, nil
}
