package controllers

import (
	"app/config"
	"app/responses"
	"app/services"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentRequest struct {
	UserId  string `json:"user_id"`
	VideoId string `json:"video_id"`
	Content string `json:"content"`
}

func Comment(c *gin.Context) {
	//var comentRequest CommentRequest
	//if err := c.ShouldBindJSON(&comentRequest); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"status_code": 1,
	//		"status_msg":  "获取数据错误",
	//	})
	//	fmt.Println("获取数据错误")
	//	return
	//}
	//content := comentRequest.Content
	//userId, err := strconv.Atoi(comentRequest.UserId)
	//videoId, err := strconv.Atoi(comentRequest.VideoId)
	content := c.DefaultQuery("content", "")
	userId, err := strconv.Atoi(c.DefaultQuery("user_id", ""))
	videoId, err := strconv.Atoi(c.DefaultQuery("video_id", ""))
	if err != nil || userId <= 0 || videoId <= 0 {
		c.JSON(http.StatusBadRequest, responses.CommentResponse{
			StatusCode: 1,
			StatusMsg:  "无效id",
		})
		return
	}
	// 验证视频是否存在
	if err = services.FindVideo(videoId, c); err != nil {
		return
	}
	if err = services.Comment(userId, videoId, content, c); err != nil {
		return
	}
	return
}

type CommentListRequest struct {
	VideoId string `json:"video_id"`
}

func List(c *gin.Context) {
	// 验证video_id
	videoId := c.DefaultQuery("video_id", "0")
	videoIdInt, err := strconv.Atoi(videoId)
	if err != nil || videoIdInt < 1 {
		c.JSON(http.StatusBadRequest, responses.CommentListResponse{
			StatusCode: 1,
			StatusMsg:  "Invalid Video ID.",
		})
		return
	}
	// 验证视频是否存在
	if err = services.FindVideo(videoIdInt, c); err != nil {
		return
	}

	// 从Redis缓存中查询评论列表
	// 获取视频的评论数据，按时间从旧到新排序
	comments, err := config.Rdb.ZRange(context.Background(), "video_comments:"+strconv.Itoa(videoIdInt), 0, -1).Result()
	if err != nil {
		// 处理从 Redis 获取评论数据失败的情况
		fmt.Println("从 Redis 获取评论数据失败:", err)
	}

	// 遍历评论数据
	for _, comment := range comments {
		// 处理每条评论数据
		fmt.Println(comment)
	}
	var commentListResponses []responses.CommentResItem
	for _, comment := range comments {
		commentListResponses = append(commentListResponses, responses.CommentResItem{
			Content: comment,
		})
	}

	c.JSON(http.StatusOK, responses.CommentListResponse{
		StatusCode:  0,
		StatusMsg:   "Success",
		CommentList: &commentListResponses,
	})
}
