package model

import (
	"github.com/zhufuyi/sponge/pkg/mysql"
)

// PostLatest 最新帖子
type PostLatest struct {
	mysql.Model `gorm:"embedded"` // embed id and time

	PostID  uint64 `gorm:"column:post_id;type:bigint(20) unsigned;NOT NULL" json:"postId"` // 帖子id
	UserID  uint64 `gorm:"column:user_id;type:bigint(20) unsigned;NOT NULL" json:"userId"` // 用户id
	DelFlag int    `gorm:"column:del_flag;type:tinyint(4);NOT NULL" json:"delFlag"`        // 删除方式，0:正常, 1:用户删除, 2:管理员删除
}

// TableName table name
func (m *PostLatest) TableName() string {
	return "post_latest"
}
