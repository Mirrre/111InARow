package services

import (
	"app/config"
	"app/consts"
	"app/models"
	"app/responses"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

var ctx = context.Background()

func FindUser(userId int, c *gin.Context) (err error) {
	db := c.MustGet("db").(*gorm.DB)

	// 验证 to_user_id 是一个存在的用户
	var toUser models.User
	if err = db.First(&toUser, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, responses.CommentResponse{
			StatusCode: 1,
			StatusMsg:  "无法找到to_user_id",
		})
		return err
	}
	return nil
}
func Follow(fromUserId int, toUserId int, c *gin.Context) (err error) {
	db := c.MustGet("db").(*gorm.DB)
	relation := models.Relation{
		FromUserId: fromUserId,
		ToUserId:   toUserId,
		CreatedAt:  time.Time{},
	}
	if err = db.Create(&relation).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) { // 重复关注
			c.JSON(http.StatusBadRequest, responses.CommentResponse{
				StatusCode: 1,
				StatusMsg:  "你已经关注过该用户了",
			})
			fmt.Println("你已经关注过该用户了")
			return
		}
		// 其它创建失败错误
		c.JSON(http.StatusInternalServerError, responses.CommentResponse{
			StatusCode: 1,
			StatusMsg:  "Failed to follow.",
		})
		return
	}
	//在Redis中创建关注数据
	err = config.Rdb.SAdd(ctx, fmt.Sprintf(consts.CacheKeyFollow, fromUserId), toUserId).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.CommentResponse{
			StatusCode: 1,
			StatusMsg:  "Failed to follow",
		})
		return
	}
	err = config.Rdb.SAdd(ctx, fmt.Sprintf(consts.CacheKeyFollower, toUserId), fromUserId).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.CommentResponse{
			StatusCode: 1,
			StatusMsg:  "Failed to follower",
		})
		return
	}
	return nil
}

func UnFollow(fromUserId int, toUserId int, c *gin.Context) (err error) {
	db := c.MustGet("db").(*gorm.DB)
	relationToDelete := models.Relation{
		FromUserId: fromUserId,
		ToUserId:   toUserId,
	}
	tx := db.Where("from_user_id = ? and to_user_id = ?", fromUserId, toUserId).
		Delete(&relationToDelete)
	if tx.RowsAffected == 0 { // 删除了0条记录，说明这条关注关系不存在
		c.JSON(http.StatusNotFound, responses.CommentResponse{
			StatusCode: 1,
			StatusMsg:  "You haven't followed this user.",
		})
		err = errors.New("你还没有关注此用户")
		return
	}
	if tx.Error != nil { // 其它删除失败错误
		c.JSON(http.StatusInternalServerError, responses.CommentResponse{
			StatusCode: 1,
			StatusMsg:  "Failed to unfollow.",
		})
		err = errors.New("无法取关")
		return
	}
	//在Redis中取消关注数据
	err = config.Rdb.SRem(ctx, fmt.Sprintf(consts.CacheKeyFollow, fromUserId), toUserId).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.CommentResponse{
			StatusCode: 1,
			StatusMsg:  "Failed to un-follow",
		})
		return
	}
	err = config.Rdb.SRem(ctx, fmt.Sprintf(consts.CacheKeyFollower, toUserId), fromUserId).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.CommentResponse{
			StatusCode: 1,
			StatusMsg:  "Failed to un-follower",
		})
		return
	}
	return nil
}
