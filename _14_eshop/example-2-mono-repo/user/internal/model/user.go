package model

import (
	"github.com/zhufuyi/sponge/pkg/ggorm"
)

type User struct {
	ggorm.Model `gorm:"embedded"` // embed id and time

	Name      string `gorm:"column:name;type:char(50);NOT NULL" json:"name"`                   // 用户名
	Password  string `gorm:"column:password;type:char(100);NOT NULL" json:"password"`          // 密码
	Email     string `gorm:"column:email;type:char(50);NOT NULL" json:"email"`                 // 邮箱
	Phone     string `gorm:"column:phone;type:char(30);NOT NULL" json:"phone"`                 // 电话号码
	Avatar    string `gorm:"column:avatar;type:varchar(200)" json:"avatar"`                    // 头像
	Age       int    `gorm:"column:age;type:tinyint(4);NOT NULL" json:"age"`                   // age
	Gender    int    `gorm:"column:gender;type:tinyint(4);NOT NULL" json:"gender"`             // 性别, 1:男, 2:女, 其他值:未知
	Status    int    `gorm:"column:status;type:tinyint(4);NOT NULL" json:"status"`             // 账号状态, 0:待验证, 1:正常, 2:封禁
	LoginAt   uint64 `gorm:"column:login_at;type:bigint(20) unsigned;NOT NULL" json:"loginAt"` // 登录时间戳
	LoginType int    `gorm:"column:login_type;type:tinyint(4);NOT NULL" json:"loginType"`      // 登录类型, 1:web, 2:mobile, 3:desktop, 4:api
}

// TableName table name
func (m *User) TableName() string {
	return "user"
}
