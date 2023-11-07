package services

import (
	"app/config"
	"app/consts"
	"app/models"
	"app/responses"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

func FindVideo(videoId int, c *gin.Context) (err error) {
	db := c.MustGet("db").(*gorm.DB)
	//验证视频是否存在
	var video models.Video
	if err = db.First(&video, videoId).Error; err != nil {
		c.JSON(http.StatusNotFound, responses.CommentResponse{
			StatusCode: 1,
			StatusMsg:  "未找到该视频",
		})
		fmt.Println("未找到该视频")
		return
	}
	return nil
}

func Comment(userId int, videoId int, content string, c *gin.Context) (err error) {
	db := c.MustGet("db").(*gorm.DB)
	// 验证评论长度
	if len(content) == 0 {
		c.JSON(http.StatusBadRequest, responses.CommentResponse{
			StatusCode: 1,
			StatusMsg:  "不允许发布空评论",
		})
		fmt.Println("不允许发布空评论")
		return
	}
	if len(content) > consts.MaxCommentLength {
		c.JSON(http.StatusBadRequest, responses.CommentResponse{
			StatusCode: 1,
			StatusMsg:  "评论内容过长",
		})
		fmt.Println("评论内容过长")
		return
	}

	// 在数据库中创建评论
	comment := models.Comment{
		UserId:    userId,
		VideoId:   videoId,
		Content:   content,
		CreatedAt: time.Now(),
	}
	result := db.Create(&comment)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, responses.CommentResponse{
			StatusCode: 1,
			StatusMsg:  "创建评论失败",
		})
		fmt.Println("创建评论失败")
		return
	}

	c.JSON(http.StatusOK, responses.CommentResponse{
		StatusCode: 0,
		StatusMsg:  "评论成功",
	})

	// 将数据存入redis
	z := redis.Z{
		Score:  float64(comment.CreatedAt.Unix()), // 使用评论的创建时间作为分数（Score）
		Member: comment.Content,                   // 使用评论内容作为成员（Member）
	}

	err = config.Rdb.ZAdd(ctx, "video_comments:"+strconv.Itoa(videoId), z).Err()
	if err != nil {
		// 处理存储到 Redis 失败的情况
		fmt.Println("存储评论到 Redis 失败:", err)
		return err
	}
	return nil
}

//
//func ClearCommentListCache(c *gin.Context, videoId int) {
//	// 清空之前的缓存
//	err := config.Rdb.Del(context.Background(), fmt.Sprintf(consts.CacheKeyComments, videoId)).Err()
//	if err != nil {
//		fmt.Printf("Failed to clear comment list cache: %v\n", err)
//	}
//
//	// 重新缓存评论列表
//	go cacheCommentList(c, videoId)
//}

// CacheCommentList 缓存评论列表
func CacheCommentList(c *gin.Context, videoId int) {
	db := c.MustGet("db").(*gorm.DB)

	var comments []models.Comment
	err := db.Where("video_id = ?", videoId).Find(&comments).Error
	if err != nil {
		fmt.Printf("Failed to fetch comments: %v\n", err)
		return
	}

	commentItems := make([]*responses.CommentResItem, 0, len(comments))
	for _, comment := range comments {
		commentItem := &responses.CommentResItem{
			Id:         comment.ID,
			Content:    comment.Content,
			CreateDate: comment.CreatedAt,
		}
		commentItems = append(commentItems, commentItem)
	}

	commentListBytes, err := json.Marshal(commentItems)
	if err != nil {
		fmt.Printf("Failed to marshal comment list: %v\n", err)
		return
	}

	err = config.Rdb.Set(context.Background(), fmt.Sprintf(consts.CacheKeyComments, videoId), commentListBytes, -1).Err()
	if err != nil {
		fmt.Printf("Failed to cache comment list: %v\n", err)
	}
}
