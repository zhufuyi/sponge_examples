package model

import (
	"github.com/zhufuyi/sponge/pkg/mysql"
)

// Course 课程
type Course struct {
	mysql.Model `gorm:"embedded"`

	Code     string `gorm:"column:code;type:varchar(10);NOT NULL" json:"code"`         // 课程代号
	Name     string `gorm:"column:name;type:varchar(50);NOT NULL" json:"name"`         // 课程名称
	Credit   int    `gorm:"column:credit;type:tinyint(4);NOT NULL" json:"credit"`      // 学分
	College  string `gorm:"column:college;type:varchar(50);NOT NULL" json:"college"`   // 学院
	Semester string `gorm:"column:semester;type:varchar(20);NOT NULL" json:"semester"` // 学期
	Time     string `gorm:"column:time;type:varchar(30);NOT NULL" json:"time"`         // 上课时间
	Place    string `gorm:"column:place;type:varchar(30);NOT NULL" json:"place"`       // 上课地点
}

// TableName table name
func (m *Course) TableName() string {
	return "course"
}
