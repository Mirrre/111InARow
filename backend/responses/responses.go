package responses

import "time"

// LoginResponse 登录返回响应结构体
type LoginResponse struct {
	// 状态码，0-成功，其他值-失败
	StatusCode int `json:"status_code"`
	// 返回状态描述
	StatusMsg string `json:"status_msg"`
	// 用户鉴权token
	Token string `json:"token"`
	// 用户id
	UserID uint `json:"user_id"`
}

// RegisterResponse 注册返回响应结构体
type RegisterResponse struct {
	// 状态码，0-成功，其他值-失败
	StatusCode int `json:"status_code"`
	// 返回状态描述
	StatusMsg string `json:"status_msg"`
	// 用户鉴权token
	Token string `json:"token"`
	// 用户id
	UserID uint `json:"user_id"`
}

// UserResponse 用户信息返回响应结构体
type UserResponse struct {
	// 状态码，0-成功，其他值-失败
	StatusCode int `json:"status_code"`
	// 返回状态描述
	StatusMsg string `json:"status_msg"`
	// 用户信息
	User User `json:"user"`
}

// User 用户信息结构体
type User struct {
	// 用户id
	ID uint `json:"id"`
	// 用户名称
	Name string `json:"name"`
	// 用户头像
	Avatar string `json:"avatar"`
	// 用户个人页顶部大图
	BackgroundImage string `json:"background_image"`
	// 个人简介
	Signature string `json:"signature"`

	// 关注总数
	FollowCount int `json:"follow_count"`
	// 粉丝总数
	FollowerCount int `json:"follower_count"`
	// 获赞数量
	TotalFavorited int `json:"total_favorited"`
	// 作品数
	WorkCount int `json:"work_count"`
	// 喜欢数
	FavoriteCount int `json:"favorite_count"`

	// true-已关注，false-未关注
	IsFollowed bool `json:"is_follow"`
}

// VideoSearchResponse 视频返回响应结构体
type VideoSearchResponse struct {
	// 状态码，0-成功，其他值-失败
	StatusCode int64 `json:"status_code"`
	// 返回状态描述
	StatusMsg string `json:"status_msg"`
	// 视频列表
	VideoList []Video `json:"video_list"`
}

// Video 视频信息结构体
type Video struct {
	// 视频唯一标识
	Id int `json:"id"`
	// 作者id
	UserId int `json:"user_id"`
	// 作者名称
	UserName string `json:"username"`
	// 作者头像
	Avatar string `json:"avatar"`
	// 封面地址
	Thumb string `json:"thumb"`
	// 视频地址
	VideoUrl string `json:"video_url"`
	// 发布时间
	Mtime string `json:"mtime"`
	// 绰号
	NickName string `json:"nickName"`
	// 分享的数量
	ShareNum int `json:"share_num" gorm:"default:0;not null"`
	// 收藏的数量
	CollectNum int `json:"collect_num" gorm:"default:0;not null"`
	// 评论的数量
	CommentNum int `json:"comment_num" gorm:"default:0;not null"`
	// 喜欢的数量
	LikeNum int `json:"like_num" gorm:"default:0;not null"`
	// 视频标题
	Title string `json:"title"`
	// 是否点赞
	IsFollow string `json:"is_follow"`
	// 是否收藏
	IsFollow2 string `json:"is_follow2"`
	// 是否关注
	IsFollow3 string `json:"is_follow3"`
}

type CommentResponse struct {
	StatusCode int
	StatusMsg  string
}

type CommentResItem struct {
	Id         int       `json:"id"`
	Content    string    `json:"content"`
	CreateDate time.Time `json:"create_date"`
}
