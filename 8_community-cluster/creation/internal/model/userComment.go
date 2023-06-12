package model

import (
	"github.com/zhufuyi/sponge/pkg/mysql"
)

// UserComment 用户评论
type UserComment struct {
	mysql.Model `gorm:"embedded"` // embed id and time

	CommentID uint64 `gorm:"column:comment_id;type:bigint(20) unsigned;NOT NULL" json:"commentId"` // 评论id
	UserID    uint64 `gorm:"column:user_id;type:bigint(20) unsigned;NOT NULL" json:"userId"`       // 用户id
	DelFlag   int    `gorm:"column:del_flag;type:tinyint(4);NOT NULL" json:"delFlag"`              // 删除方式，0:正常, 1:用户删除, 2:管理员删除
}

// TableName table name
func (m *UserComment) TableName() string {
	return "user_comment"
}
