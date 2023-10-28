package models

import (
	"github.com/u2takey/go-utils/rand"
	"gorm.io/gorm"
	"time"
)

func init() {
	// 使用当前时间的纳秒数来初始化随机数生成器
	rand.Seed(time.Now().UnixNano())
}

func getRandomIndex() int {
	// 使用随机索引选择切片中的一个字符串
	return rand.Intn(6)
}

var AvatarUrls = []string{
	"https://c-ssl.duitang.com/uploads/item/202102/04/20210204151116_snwek.jpg",
	"https://c-ssl.duitang.com/uploads/item/202102/04/20210204151117_ufhjb.jpg",
	"https://c-ssl.duitang.com/uploads/item/202102/04/20210204151118_ytkwe.jpg",
	"https://c-ssl.duitang.com/uploads/item/202102/04/20210204151119_mtxcj.jpg",
	"https://c-ssl.duitang.com/uploads/item/202102/04/20210204151120_xxfqi.jpg",
	"https://c-ssl.duitang.com/uploads/item/202102/04/20210204151123_jungh.jpg",
}

var BackgroundUrls = []string{
	"https://c-ssl.duitang.com/uploads/blog/202205/16/20220516164359_e2a6f.jpg",
	"https://c-ssl.duitang.com/uploads/blog/202205/15/20220515195808_2491f.jpeg",
	"https://c-ssl.duitang.com/uploads/blog/202205/15/20220515195807_45e27.jpeg",
	"https://c-ssl.duitang.com/uploads/blog/202207/28/20220728175554_dc1f0.jpeg",
	"https://c-ssl.duitang.com/uploads/blog/202207/28/20220728175555_8db48.jpeg",
	"https://c-ssl.duitang.com/uploads/blog/202207/28/20220728175600_e25d2.jpeg",
}

var Signatures = []string{
	"我相信人生最美的风景不在远方，而是在每一个细微的瞬间，于平凡中发现不凡，于日常中感受奇迹，用心灵的眼睛欣赏这个世界的美丽。",
	"在这个喧嚣的世界里，我选择保持内心的宁静，追求自我成长和心灵的宽广，成为那个无论风雨如何，都能散发光芒的人。",
	"生命如同一本书，我选择用勇气和冒险的笔触书写，用热情和坚持的章节编织，让每一页都闪耀着真实、勇敢和无限可能。",
	"我相信人生的真正意义不在于追逐金钱和功名，而是在于散播爱与善良，用自己的存在为世界带来一丝温暖和美好。",
	"我视挫折为成长的阶梯，相信坚持和努力的力量，每一次跌倒都是为了更坚定地站起来，继续向前，追寻心中远大的梦想。",
	"在这个瞬息万变的时代，我选择保持真实，与自己的内心对话，倾听内在的声音，成为一个与自己和谐相处的灵魂旅者。",
}

// User 表示应用中的用户
type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Profile  UserProfile `gorm:"foreignKey:UserID"`
}

// UserProfile 表示用户的额外信息
type UserProfile struct {
	gorm.Model
	UserID         uint
	Avatar         string //头像
	Background     string //背景
	Signature      string //个人简介
	FollowCount    int    `gorm:"default:0"` // 关注总数
	FollowerCount  int    `gorm:"default:0"` // 粉丝总数
	TotalFavorited int    `gorm:"default:0"` // 获赞数量
	WorkCount      int    `gorm:"default:0"` // 作品数
	FavoriteCount  int    `gorm:"default:0"` // 喜欢数
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	userProfile := UserProfile{
		UserID:         u.ID,
		Avatar:         AvatarUrls[getRandomIndex()],
		Background:     BackgroundUrls[getRandomIndex()],
		Signature:      Signatures[getRandomIndex()],
		FollowCount:    0,
		FollowerCount:  0,
		TotalFavorited: 0,
		WorkCount:      0,
		FavoriteCount:  0,
	}
	if err = tx.Create(&userProfile).Error; err != nil {
		return err
	}
	return nil
}
