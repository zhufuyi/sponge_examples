package model

import (
	"github.com/zhufuyi/sponge/pkg/mysql"
)

// CommentHot 热门评论
type CommentHot struct {
	mysql.Model `gorm:"embedded"` // embed id and time

	UserID    uint64 `gorm:"column:user_id;type:bigint(20) unsigned;NOT NULL" json:"user_id"`       // 用户id
	CommentID uint64 `gorm:"column:comment_id;type:bigint(20) unsigned;NOT NULL" json:"comment_id"` // 评论id
	PostID    uint64 `gorm:"column:post_id;type:bigint(20) unsigned;NOT NULL" json:"post_id"`       // 帖子id
	ParentID  uint64 `gorm:"column:parent_id;type:bigint(20) unsigned;NOT NULL" json:"parent_id"`   // 父id
	Score     int    `gorm:"column:score;type:tinyint(4);NOT NULL" json:"score"`                    // 分数
	DelFlag   int    `gorm:"column:del_flag;type:tinyint(4);NOT NULL" json:"del_flag"`              // 删除方式，0:正常, 1:用户删除, 2:管理员删除
}

// TableName table name
func (m *CommentHot) TableName() string {
	return "comment_hot"
}
