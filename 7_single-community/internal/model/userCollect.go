package model

import (
	"github.com/zhufuyi/sponge/pkg/mysql"
)

// UserCollect 用户收藏
type UserCollect struct {
	mysql.Model `gorm:"embedded"` // embed id and time

	UserID uint64 `gorm:"column:user_id;type:bigint(20) unsigned;NOT NULL" json:"user_id"` // 用户id
	PostID uint64 `gorm:"column:post_id;type:bigint(20) unsigned;NOT NULL" json:"post_id"` // 贴子id
}

// TableName table name
func (m *UserCollect) TableName() string {
	return "user_collect"
}
