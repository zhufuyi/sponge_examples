// Package types define the structure of request parameters and respond results in this package
package types

import (
	"time"

	"github.com/zhufuyi/sponge/pkg/mysql/query"
)

var _ time.Time

// Tip: suggested filling in the binding rules https://github.com/go-playground/validator in request struct fields tag.

// CreateTeacherRequest request params
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

// UpdateTeacherByIDRequest request params
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

// TeacherObjDetail detail
type TeacherObjDetail struct {
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

// CreateTeacherRespond only for api docs
type CreateTeacherRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		ID uint64 `json:"id"` // id
	} `json:"data"` // return data
}

// UpdateTeacherByIDRespond only for api docs
type UpdateTeacherByIDRespond struct {
	Result
}

// GetTeacherByIDRespond only for api docs
type GetTeacherByIDRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Teacher TeacherObjDetail `json:"teacher"`
	} `json:"data"` // return data
}

// DeleteTeacherByIDRespond only for api docs
type DeleteTeacherByIDRespond struct {
	Result
}

// DeleteTeachersByIDsRequest request params
type DeleteTeachersByIDsRequest struct {
	IDs []uint64 `json:"ids" binding:"min=1"` // id list
}

// DeleteTeachersByIDsRespond only for api docs
type DeleteTeachersByIDsRespond struct {
	Result
}

// GetTeacherByConditionRequest request params
type GetTeacherByConditionRequest struct {
	query.Conditions
}

// GetTeacherByConditionRespond only for api docs
type GetTeacherByConditionRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Teacher TeacherObjDetail `json:"teacher"`
	} `json:"data"` // return data
}

// ListTeachersByIDsRequest request params
type ListTeachersByIDsRequest struct {
	IDs []uint64 `json:"ids" binding:"min=1"` // id list
}

// ListTeachersByIDsRespond only for api docs
type ListTeachersByIDsRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Teachers []TeacherObjDetail `json:"teachers"`
	} `json:"data"` // return data
}

// ListTeachersRequest request params
type ListTeachersRequest struct {
	query.Params
}

// ListTeachersRespond only for api docs
type ListTeachersRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Teachers []TeacherObjDetail `json:"teachers"`
	} `json:"data"` // return data
}
