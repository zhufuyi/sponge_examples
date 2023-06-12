package model

import (
	"github.com/zhufuyi/sponge/pkg/mysql"
)

// UserFollowing 用户粉丝
type UserFollowing struct {
	mysql.Model `gorm:"embedded"` // embed id and time

	UserID      uint64 `gorm:"column:user_id;type:bigint(20) unsigned;NOT NULL" json:"userId"`           // 用户id
	FollowedUid uint64 `gorm:"column:followed_uid;type:bigint(20) unsigned;NOT NULL" json:"followedUid"` // 关注用户id
	Status      int    `gorm:"column:status;type:tinyint(4);NOT NULL" json:"status"`                     // 关注状态，0:删除, 1:正常
}

// TableName table name
func (m *UserFollowing) TableName() string {
	return "user_following"
}
