package models

import (
	"gorm.io/gorm"
	"time"
)

type Video struct {
	gorm.Model
	UserID        uint      `gorm:"index:idx_user_created" json:"user_id"`                             // 作者id
	User          User      `gorm:"foreignKey:UserID"`                                                 // 作者信息
	Title         string    `json:"title"`                                                             // 视频标题
	PlayUrl       string    `json:"play_url"`                                                          // 视频地址
	CoverUrl      string    `json:"cover_url"`                                                         // 视频封面地址
	FavoriteCount uint      `gorm:"default:0;not null" json:"favorite_count"`                          // 点赞人数
	CommentCount  uint      `gorm:"default:0;not null" json:"comment_count"`                           // 评论人数
	CollectCount  uint      `gorm:"default:0;not null" json:"collect_count"`                           // 收藏人数
	PublishTime   time.Time `gorm:"index:idx_publish_time;index:idx_user_created" json:"published_at"` // 发布时间
}

// AfterCreate hook for the Video model.
func (video *Video) AfterCreate(tx *gorm.DB) (err error) {
	// 发布者的作品数 + 1
	err = tx.Model(&UserProfile{}).Where("user_id = ?", video.UserID).
		UpdateColumn("work_count", gorm.Expr("work_count + 1")).Error
	if err != nil {
		return err
	}
	return nil
}

// AfterDelete hook for the Video model.
func (video *Video) AfterDelete(tx *gorm.DB) (err error) {
	// 发布者的作品数 - 1
	err = tx.Model(&UserProfile{}).Where("user_id = ?", video.UserID).
		UpdateColumn("work_count", gorm.Expr(
			"CASE WHEN work_count > 0 THEN work_count - 1 ELSE 0 END")).Error
	if err != nil {
		return err
	}
	return nil
}
