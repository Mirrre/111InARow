package controllers

import (
	"app/models"
	"app/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"os"
	"strconv"
	"time"
)

type PublishRequest struct {
	Title  string `json:"title"`
	UserId string `json:"user_id"`
}

func Publish(c *gin.Context) {
	//token := c.PostForm("token")
	//title := c.PostForm("title")
	request := PublishRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": 1,
			"status_msg":  "获取数据错误",
		})
		fmt.Println("获取数据错误")
		return
	}
	userId, err := strconv.Atoi(request.UserId)
	title := request.Title
	//userid, err := utils.ValidateToken(token)
	//userId := int(userid)

	// 获取上传视频
	videoFile, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": 1,
			"status_msg":  "上传视频数据错误",
		})
		fmt.Println("上传视频数据错误")
		return
	}
	// 将视频存在本地
	// 设置本地视频路径
	videoLocalPath := fmt.Sprintf("./uploads/video/%s.mp4", title)
	err = c.SaveUploadedFile(videoFile, videoLocalPath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": 1,
			"status_msg":  "视频保存本地错误",
		})
		fmt.Println("视频保存本地错误")
		return
	}
	// 自定义的视频文件名
	key := fmt.Sprintf("video/%d/%s.mp4", userId, title)
	videoUrl, err := services.UploadVideo(videoLocalPath, key)
	if err == nil {
		fmt.Println(videoUrl)
		// 通过本地视频文件生成本地封面截图
		coverLocalPath := services.GetSnapshot(videoUrl, "./uploads/cover/", title, 1)
		// 自定义的截图文件名
		key = fmt.Sprintf("cover/%d/%s.jpeg", userId, title)
		coverUrl, err := services.UploadCover(coverLocalPath, key)
		// 删除本地文件
		err = os.Remove("./uploads/video/" + title + ".mp4")
		err = os.Remove("./uploads/cover/" + title + ".jpeg")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status_code": 1,
				"status_msg":  "无法删除本地文件",
			})
			fmt.Println("无法删除本地文件")
			return
		}
		if err == nil {
			db := c.MustGet("db").(*gorm.DB)
			var user models.User
			result := db.First(&user, "id = ?", userId)
			if result.Error != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status_code": 1,
					"status_msg":  "获取用户头像错误",
				})
				fmt.Println("获取用户头像错误")
				return
			}

			newVideo := models.Video{
				Avatar:   user.Avatar,
				UserId:   userId,
				Thumb:    coverUrl,
				VideoUrl: videoUrl,
				Title:    title,
				Mtime:    time.Now().String(),
			}
			db.Create(&newVideo)
			c.JSON(http.StatusOK, gin.H{
				"status_code": 0,
				"status_msg":  "视频上传成功",
			})

		}
	}
}

//type SearchRequest struct {
//	Title string
//}

//func Search(c *gin.Context) {
//	var searchRequest SearchRequest
//	if err := c.ShouldBindJSON(&searchRequest); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"status_code": 1,
//			"status_msg":  "获取数据错误",
//		})
//		fmt.Println("获取数据错误")
//		return
//	}
//	title := searchRequest.Title
//}
