package model

import (
	"github.com/zhufuyi/sponge/pkg/mysql"
)

// CommentContent 评论内容
type CommentContent struct {
	mysql.Model `gorm:"embedded"` // embed id and time

	CommentID  uint64 `gorm:"column:comment_id;type:bigint(20) unsigned;NOT NULL" json:"comment_id"` // 评论id
	Content    string `gorm:"column:content;type:text;NOT NULL" json:"content"`                      // 评论内容
	DeviceType string `gorm:"column:device_type;type:varchar(50);NOT NULL" json:"device_type"`       // 设备类型
	IP         string `gorm:"column:ip;type:char(16);NOT NULL" json:"ip"`                            // ip地址
}

// TableName table name
func (m *CommentContent) TableName() string {
	return "comment_content"
}
