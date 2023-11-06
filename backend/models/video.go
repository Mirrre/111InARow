package models

import (
	"gorm.io/gorm"
)

type Video struct {
	Id     int `json:"id"`
	UserID int `gorm:"index:idx_user_created" json:"user_id"` // 作者id
	//User     User   `gorm:"foreignKey:UserID"`                     // 作者信息
	UserName string `json:"username"`
	//Title         string    `json:"title"`                                                             // 视频标题
	//PlayUrl       string    `json:"play_url"`                                                          // 视频地址
	//CoverUrl      string    `json:"cover_url"`                                                         // 视频封面地址
	//FavoriteCount int      `gorm:"default:0;not null" json:"favorite_count"`                          // 点赞人数
	//CommentCount  int      `gorm:"default:0;not null" json:"comment_count"`                           // 评论人数
	//CollectCount  int      `gorm:"default:0;not null" json:"collect_count"`                           // 收藏人数
	//PublishTime   time.Time `gorm:"index:idx_publish_time;index:idx_user_created" json:"published_at"` // 发布时间
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
	err = tx.Model(&User{}).Where("id = ?", video.UserID).
		UpdateColumn("work_count", gorm.Expr("work_count + 1")).Error
	if err != nil {
		return err
	}
	return nil
}

// AfterDelete hook for the Video model.
func (video *Video) AfterDelete(tx *gorm.DB) (err error) {
	// 发布者的作品数 - 1
	err = tx.Model(&User{}).Where("id = ?", video.UserID).
		UpdateColumn("work_count", gorm.Expr(
			"CASE WHEN work_count > 0 THEN work_count - 1 ELSE 0 END")).Error
	if err != nil {
		return err
	}
	return nil
}
