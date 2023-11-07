package models

import (
	"gorm.io/gorm"
)

type Video struct {
	Id         int    `json:"id"`
	UserId     int    `gorm:"index:idx_user_created" json:"user_id"` // 作者id
	UserName   string `json:"username" gorm:"column:username"`
	Avatar     string `json:"avatar"`
	Thumb      string `json:"thumb"`
	VideoUrl   string `json:"video_url"`
	Mtime      string `json:"mtime"`
	NickName   string `json:"nickName"`
	ShareNum   int    `json:"share_num" gorm:"default:0;not null"`
	CollectNum int    `json:"collect_num" gorm:"default:0;not null"`
	CommentNum int    `json:"comment_num" gorm:"default:0;not null"`
	LikeNum    int    `json:"like_num" gorm:"default:0;not null"`
	Title      string `json:"title"`

	IsFollow  string `json:"is_follow"`
	IsFollow2 string `json:"is_follow2"`
}

// AfterCreate hook for the Video model.
func (video *Video) AfterCreate(tx *gorm.DB) (err error) {
	// 发布者的作品数 + 1
	err = tx.Model(&User{}).Where("id = ?", video.UserId).
		UpdateColumn("work_count", gorm.Expr("work_count + 1")).Error
	if err != nil {
		return err
	}
	return nil
}

// AfterDelete hook for the Video model.
func (video *Video) AfterDelete(tx *gorm.DB) (err error) {
	// 发布者的作品数 - 1
	err = tx.Model(&User{}).Where("id = ?", video.UserId).
		UpdateColumn("work_count", gorm.Expr(
			"CASE WHEN work_count > 0 THEN work_count - 1 ELSE 0 END")).Error
	if err != nil {
		return err
	}
	return nil
}
