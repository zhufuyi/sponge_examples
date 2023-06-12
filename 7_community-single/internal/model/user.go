package model

import (
	"github.com/zhufuyi/sponge/pkg/mysql"
	"time"
)

// User 用户信息
type User struct {
	mysql.Model `gorm:"embedded"` // embed id and time

	Name     string    `gorm:"column:name;type:varchar(50);NOT NULL" json:"name"`          // 用户名
	NickName string    `gorm:"column:nick_name;type:varchar(50);NOT NULL" json:"nickName"` // 用户昵称
	Password string    `gorm:"column:password;type:varchar(100);NOT NULL" json:"password"` // 密码
	Email    string    `gorm:"column:email;type:varchar(50);NOT NULL" json:"email"`        // 邮件
	Phone    string    `gorm:"column:phone;type:varchar(30);NOT NULL" json:"phone"`        // 手机号码
	Avatar   string    `gorm:"column:avatar;type:varchar(200)" json:"avatar"`              // 头像
	Gender   int       `gorm:"column:gender;type:tinyint(4);NOT NULL" json:"gender"`       // 性别，1:男，2:女，其他值:未知
	Age      int       `gorm:"column:age;type:tinyint(4);NOT NULL" json:"age"`             // 年龄
	Birthday string    `gorm:"column:birthday;type:varchar(30);NOT NULL" json:"birthday"`  // 出生日期
	LoginAt  time.Time `gorm:"column:login_at;type:datetime" json:"loginAt"`               // 登录时间
	LoginIP  string    `gorm:"column:login_ip;type:char(16)" json:"loginIp"`               // 登录ip
	Status   int       `gorm:"column:status;type:tinyint(4);NOT NULL" json:"status"`       // 状态, 1:正常, 2:删除, 3:封禁
}

// TableName table name
func (m *User) TableName() string {
	return "user"
}
