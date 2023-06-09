package model

import (
	"github.com/zhufuyi/sponge/pkg/mysql"
)

// UserLike 用户点赞
type UserLike struct {
	mysql.Model `gorm:"embedded"` // embed id and time

	ObjType int    `gorm:"column:obj_type;type:tinyint(4);NOT NULL" json:"obj_type"`        // 点赞类型，0:未知, 1:帖子, 2:评论
	ObjID   uint64 `gorm:"column:obj_id;type:bigint(20) unsigned;NOT NULL" json:"obj_id"`   // 点赞类型id
	UserID  uint64 `gorm:"column:user_id;type:bigint(20) unsigned;NOT NULL" json:"user_id"` // 用户id
	Status  int    `gorm:"column:status;type:tinyint(4);NOT NULL" json:"status"`            // 点赞状态，0:未点赞, 1:已点赞
}

// TableName table name
func (m *UserLike) TableName() string {
	return "user_like"
}
