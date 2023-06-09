package model

import (
	"github.com/zhufuyi/sponge/pkg/mysql"
)

// UserPost 最新帖子
type UserPost struct {
	mysql.Model `gorm:"embedded"` // embed id and time

	PostID  uint64 `gorm:"column:post_id;type:bigint(20) unsigned;NOT NULL" json:"post_id"` // 帖子id
	UserID  uint64 `gorm:"column:user_id;type:bigint(20) unsigned;NOT NULL" json:"user_id"` // 用户id
	DelFlag int    `gorm:"column:del_flag;type:tinyint(4);NOT NULL" json:"del_flag"`        // 删除方式，0:正常, 1:用户删除, 2:管理员删除
}

// TableName table name
func (m *UserPost) TableName() string {
	return "user_post"
}
