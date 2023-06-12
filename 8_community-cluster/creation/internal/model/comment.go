package model

import (
	"github.com/zhufuyi/sponge/pkg/mysql"
)

// Comment 评论详情
type Comment struct {
	mysql.Model `gorm:"embedded"` // embed id and time

	PostID     uint64 `gorm:"column:post_id;type:bigint(20) unsigned;NOT NULL" json:"postId"`      // 帖子id
	Type       int    `gorm:"column:type;type:tinyint(4);NOT NULL" json:"type"`                    // 类型：0:未知, 1:文本, 2:图片, 3:视频
	UserID     uint64 `gorm:"column:user_id;type:bigint(20) unsigned;NOT NULL" json:"userId"`      // 用户id
	ParentID   uint64 `gorm:"column:parent_id;type:bigint(20) unsigned;NOT NULL" json:"parentId"`  // 父id
	LikeCount  uint   `gorm:"column:like_count;type:int(10) unsigned;NOT NULL" json:"likeCount"`   // 点赞数
	ReplyCount uint   `gorm:"column:reply_count;type:int(10) unsigned;NOT NULL" json:"replyCount"` // 回复数
	Score      int    `gorm:"column:score;type:tinyint(4);NOT NULL" json:"score"`                  // 分数
	ToUid      uint64 `gorm:"column:to_uid;type:bigint(20) unsigned;NOT NULL" json:"toUid"`        // 给id
	DelFlag    int    `gorm:"column:del_flag;type:tinyint(4);NOT NULL" json:"delFlag"`             // 删除方式，0:正常, 1:用户删除, 2:管理员删除
}

// TableName table name
func (m *Comment) TableName() string {
	return "comment"
}
