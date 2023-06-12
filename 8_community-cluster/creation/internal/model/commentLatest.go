package model

import (
	"github.com/zhufuyi/sponge/pkg/mysql"
)

// CommentLatest 最新评论
type CommentLatest struct {
	mysql.Model `gorm:"embedded"` // embed id and time

	CommentID uint64 `gorm:"column:comment_id;type:bigint(20) unsigned;NOT NULL" json:"commentId"` // 评论id
	PostID    uint64 `gorm:"column:post_id;type:bigint(20) unsigned;NOT NULL" json:"postId"`       // 帖子id
	ParentID  uint64 `gorm:"column:parent_id;type:bigint(20) unsigned;NOT NULL" json:"parentId"`   // 父id
	UserID    uint64 `gorm:"column:user_id;type:bigint(20) unsigned;NOT NULL" json:"userId"`       // 用户id
	Score     int    `gorm:"column:score;type:tinyint(4);NOT NULL" json:"score"`                   // 分数
	DelFlag   int    `gorm:"column:del_flag;type:tinyint(4);NOT NULL" json:"delFlag"`              // 删除方式，0:正常, 1:用户删除, 2:管理员删除
}

// TableName table name
func (m *CommentLatest) TableName() string {
	return "comment_latest"
}
