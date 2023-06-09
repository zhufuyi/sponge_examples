package model

import (
	"github.com/zhufuyi/sponge/pkg/mysql"
)

// PostHot 热门贴子
type PostHot struct {
	mysql.Model `gorm:"embedded"` // embed id and time

	UserID  uint64 `gorm:"column:user_id;type:bigint(20) unsigned;NOT NULL" json:"user_id"` // 用户id
	PostID  uint64 `gorm:"column:post_id;type:bigint(20) unsigned;NOT NULL" json:"post_id"` // 帖子id
	Score   int    `gorm:"column:score;type:tinyint(4);NOT NULL" json:"score"`              // 分数
	DelFlag int    `gorm:"column:del_flag;type:tinyint(4);NOT NULL" json:"del_flag"`        // 删除方式，0:正常, 1:用户删除, 2:管理员删除
}

// TableName table name
func (m *PostHot) TableName() string {
	return "post_hot"
}
