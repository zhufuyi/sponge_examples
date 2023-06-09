package model

import (
	"github.com/zhufuyi/sponge/pkg/mysql"
)

// Post 帖子详情
type Post struct {
	mysql.Model `gorm:"embedded"` // embed id and time

	PostType     int     `gorm:"column:post_type;type:tinyint(4);NOT NULL" json:"post_type"`               // 类型：0:未知, 1:文本, 2:图片, 3:视频
	UserID       uint64  `gorm:"column:user_id;type:bigint(20) unsigned;NOT NULL" json:"user_id"`          // 用户id
	Title        string  `gorm:"column:title;type:varchar(250);NOT NULL" json:"title"`                     // 帖子标题
	Content      string  `gorm:"column:content;type:text;NOT NULL" json:"content"`                         // 帖子内容
	ViewCount    uint    `gorm:"column:view_count;type:int(10) unsigned;NOT NULL" json:"view_count"`       // 查看帖子次数
	LikeCount    uint    `gorm:"column:like_count;type:int(10) unsigned;NOT NULL" json:"like_count"`       // 点赞数
	CommentCount uint    `gorm:"column:comment_count;type:int(10) unsigned;NOT NULL" json:"comment_count"` // 评论数
	CollectCount uint    `gorm:"column:collect_count;type:int(10) unsigned;NOT NULL" json:"collect_count"` // 收藏数
	ShareCount   uint    `gorm:"column:share_count;type:int(10) unsigned;NOT NULL" json:"share_count"`     // 分享数
	Longitude    float64 `gorm:"column:longitude;type:float;NOT NULL" json:"longitude"`                    // 经度
	Latitude     float64 `gorm:"column:latitude;type:float;NOT NULL" json:"latitude"`                      // 纬度
	Position     string  `gorm:"column:position;type:varchar(100);NOT NULL" json:"position"`               // 位置
	Visible      int     `gorm:"column:visible;type:tinyint(4);NOT NULL" json:"visible"`                   // 显示类型，0:公开,1:仅自己可见
	DelFlag      int     `gorm:"column:del_flag;type:tinyint(4);NOT NULL" json:"del_flag"`                 // 删除方式，0:正常, 1:用户删除, 2:管理员删除
}

// TableName table name
func (m *Post) TableName() string {
	return "post"
}
