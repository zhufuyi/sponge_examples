package model

import (
	"github.com/zhufuyi/sponge/pkg/mysql"
)

// Comment 评论详情
type Comment struct {
	mysql.Model `gorm:"embedded"` // embed id and time

	PostID     uint64 `gorm:"column:post_id;type:bigint(20) unsigned;NOT NULL" json:"post_id"`      // 帖子id
	Type       int    `gorm:"column:type;type:tinyint(4);NOT NULL" json:"type"`                     // 类型：0:未知, 1:文本, 2:图片, 3:视频
	UserID     uint64 `gorm:"column:user_id;type:bigint(20) unsigned;NOT NULL" json:"user_id"`      // 用户id
	ParentID   uint64 `gorm:"column:parent_id;type:bigint(20) unsigned;NOT NULL" json:"parent_id"`  // 父id
	LikeCount  uint   `gorm:"column:like_count;type:int(10) unsigned;NOT NULL" json:"like_count"`   // 点赞数
	ReplyCount uint   `gorm:"column:reply_count;type:int(10) unsigned;NOT NULL" json:"reply_count"` // 回复数
	Score      int    `gorm:"column:score;type:tinyint(4);NOT NULL" json:"score"`                   // 分数
	ToUid      uint64 `gorm:"column:to_uid;type:bigint(20) unsigned;NOT NULL" json:"to_uid"`        // 给id
	DelFlag    int    `gorm:"column:del_flag;type:tinyint(4);NOT NULL" json:"del_flag"`             // 删除方式，0:正常, 1:用户删除, 2:管理员删除
}

// TableName table name
func (m *Comment) TableName() string {
	return "comment"
}
