package controllers

//
//import (
//	"app/config"
//	"app/consts"
//	"app/models"
//	"app/responses"
//	"app/services"
//	"context"
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"gorm.io/gorm"
//	"net/http"
//	"strconv"
//)
//
//type CommentRequest struct {
//	UserId  string `json:"user_id"`
//	VideoId string `json:"video_id"`
//	Content string `json:"content"`
//}
//
//func Comment(c *gin.Context) {
//	var comentRequest CommentRequest
//	if err := c.ShouldBindJSON(&comentRequest); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"status_code": 1,
//			"status_msg":  "获取数据错误",
//		})
//		fmt.Println("获取数据错误")
//		return
//	}
//	content := comentRequest.Content
//	userId, err := strconv.Atoi(comentRequest.UserId)
//	videoId, err := strconv.Atoi(comentRequest.VideoId)
//	if err != nil || userId <= 0 || videoId <= 0 {
//		c.JSON(http.StatusBadRequest, responses.CommentResponse{
//			StatusCode: 1,
//			StatusMsg:  "无效id",
//		})
//		return
//	}
//	if err = services.FindVideo(videoId, c); err != nil {
//		return
//	}
//	if err = services.Comment(userId, videoId, content, c); err != nil {
//		return
//	}
//	return
//}
//
//func List(c *gin.Context) {
//	// 验证video_id
//	videoID := c.DefaultQuery("video_id", "0")
//	videoIDInt, err := strconv.Atoi(videoID)
//	if err != nil || videoIDInt < 1 {
//		c.JSON(http.StatusBadRequest, responses.CommentListResponse{
//			StatusCode: 1,
//			StatusMsg:  "Invalid Video ID.",
//		})
//		return
//	}
//
//	// 验证视频是否存在
//	db := c.MustGet("db").(*gorm.DB)
//	var video models.Video
//	if err := db.First(&video, videoIDInt).Error; err != nil {
//		c.JSON(http.StatusNotFound, responses.CommentResponse{
//			StatusCode: 1,
//			StatusMsg:  "Video not found.",
//		})
//		return
//	}
//
//	var commentList []models.Comment
//	result := db.Preload("User").Where("video_id = ?", videoIDInt).Find(&commentList).Order("created_at desc")
//	if result.Error != nil {
//		c.JSON(http.StatusInternalServerError, responses.CommentListResponse{
//			StatusCode: 1,
//			StatusMsg:  "Failed to fetch comments",
//		})
//		return
//	}
//
//	// 查询评论列表中有哪些评论者是当前用户关注的
//	tokenString := c.DefaultQuery("token", "")
//	currentUserID, _ := responses.ValidateToken(tokenString)
//	var commenterIDsSet = make(map[uint]bool)
//	for _, c := range commentList {
//		commenterIDsSet[c.UserID] = true
//	}
//	var commenterIDs []uint
//	for id := range commenterIDsSet {
//		commenterIDs = append(commenterIDs, id)
//	}
//	var followedIDs []uint
//	db.Table("relations").
//		Where("from_user_id = ? AND to_user_id IN ?", currentUserID, commenterIDs).
//		Pluck("to_user_id", &followedIDs)
//	var followedIDSet = make(map[uint]bool)
//	for _, id := range followedIDs {
//		followedIDSet[id] = true
//	}
//
//	// 检查Redis缓存中是否存在评论列表
//	cacheKey := fmt.Sprintf("video:%d:comments", videoIDInt)
//	commentsJSON, err := config.Rdb.Get(context.Background(), cacheKey).Result()
//	if err == nil {
//		fmt.Println("IsRedis") //Redis命中输出
//		var commentListResponses []responses.CommentResItem
//		err = json.Unmarshal([]byte(commentsJSON), &commentListResponses)
//		if err == nil {
//			c.JSON(http.StatusOK, responses.CommentListResponse{
//				StatusCode:  0,
//				StatusMsg:   "Success",
//				CommentList: &commentListResponses,
//			})
//			return
//		}
//	}
//	var commentListResponses []responses.CommentResItem
//	for _, comment := range commentList {
//		isFollowed := followedIDSet[comment.UserID]
//		commentListResponses = append(commentListResponses, responses.CommentResItem{
//			ID: comment.ID,
//			User: responses.UserResponse{
//				ID:             comment.User.ID,
//				Name:           comment.User.Username,
//				FollowCount:    comment.User.Profile.FollowCount,
//				FollowerCount:  comment.User.Profile.FollowerCount,
//				IsFollow:       isFollowed,
//				Avatar:         comment.User.Profile.Avatar,
//				Background:     comment.User.Profile.Background,
//				Signature:      comment.User.Profile.Signature,
//				TotalFavorited: comment.User.Profile.TotalFavorited,
//				WorkCount:      comment.User.Profile.WorkCount,
//				FavoriteCount:  comment.User.Profile.FavoriteCount,
//			},
//			Content:    comment.Content,
//			CreateDate: comment.CreatedAt,
//		})
//	}
//
//	// 将评论列表存入Redis缓存
//	commentsJSONByte := []byte(commentsJSON)
//	commentsJSONByte, _ = json.Marshal(commentListResponses)
//	config.Rdb.Set(context.Background(), cacheKey, commentsJSONByte, consts.CacheExpiration) // 设置适当的过期时间
//
//	c.JSON(http.StatusOK, responses.CommentListResponse{
//		StatusCode:  0,
//		StatusMsg:   "Success",
//		CommentList: &commentListResponses,
//	})
//}
