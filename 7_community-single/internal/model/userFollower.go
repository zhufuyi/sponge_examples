package model

import (
	"github.com/zhufuyi/sponge/pkg/mysql"
)

// UserFollower 用户粉丝
type UserFollower struct {
	mysql.Model `gorm:"embedded"` // embed id and time

	UserID      uint64 `gorm:"column:user_id;type:bigint(20) unsigned;NOT NULL" json:"userId"`           // 用户id
	FollowerUid uint64 `gorm:"column:follower_uid;type:bigint(20) unsigned;NOT NULL" json:"followerUid"` // 粉丝id
	Status      int    `gorm:"column:status;type:tinyint(4);NOT NULL" json:"status"`                     // 关注状态，0:删除, 1:正常
}

// TableName table name
func (m *UserFollower) TableName() string {
	return "user_follower"
}
