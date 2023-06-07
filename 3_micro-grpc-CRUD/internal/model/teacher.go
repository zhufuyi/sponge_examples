package model

import (
	"github.com/zhufuyi/sponge/pkg/mysql"
)

// Teacher 老师
type Teacher struct {
	mysql.Model `gorm:"embedded"`

	Name       string `gorm:"column:name;type:varchar(50);NOT NULL" json:"name"`              // 用户名
	Password   string `gorm:"column:password;type:varchar(100);NOT NULL" json:"password"`     // 密码
	Email      string `gorm:"column:email;type:varchar(50);NOT NULL" json:"email"`            // 邮件
	Phone      string `gorm:"column:phone;type:varchar(30);NOT NULL" json:"phone"`            // 手机号码
	Avatar     string `gorm:"column:avatar;type:varchar(200)" json:"avatar"`                  // 头像
	Gender     int    `gorm:"column:gender;type:tinyint(4);NOT NULL" json:"gender"`           // 性别，1:男，2:女，其他值:未知
	Age        int    `gorm:"column:age;type:tinyint(4);NOT NULL" json:"age"`                 // 年龄
	Birthday   string `gorm:"column:birthday;type:varchar(30);NOT NULL" json:"birthday"`      // 出生日期
	SchoolName string `gorm:"column:school_name;type:varchar(50);NOT NULL" json:"schoolName"` // 学校名称
	College    string `gorm:"column:college;type:varchar(50);NOT NULL" json:"college"`        // 学院
	Title      string `gorm:"column:title;type:varchar(10);NOT NULL" json:"title"`            // 职称
	Profile    string `gorm:"column:profile;type:text;NOT NULL" json:"profile"`               // 个人简介
}

// TableName table name
func (m *Teacher) TableName() string {
	return "teacher"
}
