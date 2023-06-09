package model

import (
	"github.com/zhufuyi/sponge/pkg/mysql"
)

// RelationNum 用户关注和粉丝数量
type RelationNum struct {
	mysql.Model `gorm:"embedded"` // embed id and time

	UserID       uint64 `gorm:"column:user_id;type:bigint(20) unsigned;NOT NULL" json:"user_id"`             // 用户id
	FollowingNum uint64 `gorm:"column:following_num;type:bigint(20) unsigned;NOT NULL" json:"following_num"` // 关注数量
	FollowerNum  uint64 `gorm:"column:follower_num;type:bigint(20) unsigned;NOT NULL" json:"follower_num"`   // 粉丝数量
}

// TableName table name
func (m *RelationNum) TableName() string {
	return "relation_num"
}
