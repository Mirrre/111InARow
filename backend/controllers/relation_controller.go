package controllers

import (
	"app/responses"
	"app/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RelationFollowRequest struct {
	FromUserId string `json:"from_user_id"`
	ToUserId   string `json:"to_user_id"`
	ActionType string `json:"action_type"`
}

func Follow(c *gin.Context) {
	relationRequest := RelationFollowRequest{}
	if err := c.ShouldBindJSON(&relationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": 1,
			"status_msg":  "获取数据错误",
		})
		fmt.Println("获取数据错误")
		return
	}
	fromUserId, err := strconv.Atoi(relationRequest.FromUserId)
	toUserId, err := strconv.Atoi(relationRequest.ToUserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.CommentResponse{
			StatusCode: 1,
			StatusMsg:  "用户id无效",
		})
		fmt.Println("用户id无效1")
		return
	}
	actionType := relationRequest.ActionType
	if err != nil || fromUserId <= 0 || toUserId <= 0 {
		c.JSON(http.StatusBadRequest, responses.CommentResponse{
			StatusCode: 1,
			StatusMsg:  "用户id无效",
		})
		fmt.Println("用户id无效2")
		return
	}

	// 验证 from_user_id 和 to_user_id 是不是同一个 id
	if toUserId == fromUserId {
		c.JSON(http.StatusBadRequest, responses.CommentResponse{
			StatusCode: 1,
			StatusMsg:  "无法关注自己",
		})
		return
	}
	// 查询用户是否存在
	err = services.FindUser(toUserId, c)
	if err != nil {
		return
	}

	// 执行关注/取关操作
	switch actionType {
	case "1": // 关注
		err = services.Follow(fromUserId, toUserId, c)
		if err != nil {
			return
		}
	case "2": // 取关
		err = services.UnFollow(fromUserId, toUserId, c)
		if err != nil {
			return
		}

	default: // 错误的 action_type
		c.JSON(http.StatusBadRequest, responses.CommentResponse{
			StatusCode: 1,
			StatusMsg:  "Invalid action_type.",
		})
		return
	}

	c.JSON(http.StatusOK, responses.CommentResponse{
		StatusCode: 0,
		StatusMsg:  "Success",
	})
	return
}
