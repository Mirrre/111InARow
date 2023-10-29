package responses

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

// VideoResponse 视频返回响应结构体
type VideoResponse struct {
	// 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	NextTime int64 `json:"next_time"`
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
	ID int64 `json:"id"`
	// 视频作者信息
	Author User `json:"author"`
	// 视频标题
	Title string `json:"title"`

	// 视频封面地址
	CoverURL string `json:"cover_url"`
	// 视频播放地址
	PlayURL string `json:"play_url"`

	// 视频的点赞总数
	FavoriteCount int64 `json:"favorite_count"`
	// 视频的评论总数
	CommentCount int64 `json:"comment_count"`

	// true-已点赞，false-未点赞
	IsFavorite bool `json:"is_favorite"`
	// true-已收藏，false-未收藏
	IsCollect bool `json:"is_collect"`
}
