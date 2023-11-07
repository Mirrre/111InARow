package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID        int       `gorm:"primary_key;not null" json:"id"`
	UserId    int       `gorm:"not null" json:"user_id"`
	VideoId   int       `gorm:"index:idx_video_comment_created;not null" json:"video_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `gorm:"index:idx_video_comment_created"`
}

func (f *Comment) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("Video ID: ", f.VideoId)
	tx.Model(&Video{}).Where("id = ?", f.VideoId).
		UpdateColumn("comment_num", gorm.Expr("comment_num + 1"))
	return
}

func (f *Comment) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Println("Video ID: ", f.VideoId)
	tx.Model(&Video{}).Where("id = ?", f.VideoId).
		UpdateColumn("comment_num", gorm.Expr(
			"CASE WHEN comment_num > 0 THEN comment_num - 1 ELSE 0 END"))
	return
}
