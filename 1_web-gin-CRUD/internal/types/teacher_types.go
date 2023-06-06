package types

import (
	"time"

	"github.com/zhufuyi/sponge/pkg/mysql/query"
)

var _ time.Time

// CreateTeacherRequest create params
// todo fill in the binding rules https://github.com/go-playground/validator
type CreateTeacherRequest struct {
	Name       string `json:"name" binding:""`       // 用户名
	Password   string `json:"password" binding:""`   // 密码
	Email      string `json:"email" binding:""`      // 邮件
	Phone      string `json:"phone" binding:""`      // 手机号码
	Avatar     string `json:"avatar" binding:""`     // 头像
	Gender     int    `json:"gender" binding:""`     // 性别，1:男，2:女，其他值:未知
	Age        int    `json:"age" binding:""`        // 年龄
	Birthday   string `json:"birthday" binding:""`   // 出生日期
	SchoolName string `json:"schoolName" binding:""` // 学校名称
	College    string `json:"college" binding:""`    // 学院
	Title      string `json:"title" binding:""`      // 职称
	Profile    string `json:"profile" binding:""`    // 个人简介
}

// UpdateTeacherByIDRequest update params
type UpdateTeacherByIDRequest struct {
	ID uint64 `json:"id" binding:""` // uint64 id

	Name       string `json:"name" binding:""`       // 用户名
	Password   string `json:"password" binding:""`   // 密码
	Email      string `json:"email" binding:""`      // 邮件
	Phone      string `json:"phone" binding:""`      // 手机号码
	Avatar     string `json:"avatar" binding:""`     // 头像
	Gender     int    `json:"gender" binding:""`     // 性别，1:男，2:女，其他值:未知
	Age        int    `json:"age" binding:""`        // 年龄
	Birthday   string `json:"birthday" binding:""`   // 出生日期
	SchoolName string `json:"schoolName" binding:""` // 学校名称
	College    string `json:"college" binding:""`    // 学院
	Title      string `json:"title" binding:""`      // 职称
	Profile    string `json:"profile" binding:""`    // 个人简介
}

// GetTeacherByIDRespond respond detail
type GetTeacherByIDRespond struct {
	ID string `json:"id"` // convert to string id

	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Name       string    `json:"name"`       // 用户名
	Password   string    `json:"password"`   // 密码
	Email      string    `json:"email"`      // 邮件
	Phone      string    `json:"phone"`      // 手机号码
	Avatar     string    `json:"avatar"`     // 头像
	Gender     int       `json:"gender"`     // 性别，1:男，2:女，其他值:未知
	Age        int       `json:"age"`        // 年龄
	Birthday   string    `json:"birthday"`   // 出生日期
	SchoolName string    `json:"schoolName"` // 学校名称
	College    string    `json:"college"`    // 学院
	Title      string    `json:"title"`      // 职称
	Profile    string    `json:"profile"`    // 个人简介
}

// DeleteTeachersByIDsRequest request form ids
type DeleteTeachersByIDsRequest struct {
	IDs []uint64 `json:"ids" binding:"min=1"` // id list
}

// GetTeachersByIDsRequest request form ids
type GetTeachersByIDsRequest struct {
	IDs []uint64 `json:"ids" binding:"min=1"` // id list
}

// GetTeachersRequest request form params
type GetTeachersRequest struct {
	query.Params // query parameters
}

// ListTeachersRespond list data
type ListTeachersRespond []struct {
	GetTeacherByIDRespond
}
