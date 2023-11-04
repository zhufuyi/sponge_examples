// Package types define the structure of request parameters and respond results in this package
package types

import (
	"time"

	"github.com/zhufuyi/sponge/pkg/mysql/query"
)

var _ time.Time

// Tip: suggested filling in the binding rules https://github.com/go-playground/validator in request struct fields tag.

// CreateTeachRequest request params
type CreateTeachRequest struct {
	TeacherID   int64  `json:"teacherId" binding:""`   // 老师id
	TeacherName string `json:"teacherName" binding:""` // 老师名称
	CourseID    int64  `json:"courseId" binding:""`    // 课程id
	CourseName  string `json:"courseName" binding:""`  // 课程名称
	Score       string `json:"score" binding:""`       // 学生评价教学质量，5个等级：A,B,C,D,E
}

// UpdateTeachByIDRequest request params
type UpdateTeachByIDRequest struct {
	ID uint64 `json:"id" binding:""` // uint64 id

	TeacherID   int64  `json:"teacherId" binding:""`   // 老师id
	TeacherName string `json:"teacherName" binding:""` // 老师名称
	CourseID    int64  `json:"courseId" binding:""`    // 课程id
	CourseName  string `json:"courseName" binding:""`  // 课程名称
	Score       string `json:"score" binding:""`       // 学生评价教学质量，5个等级：A,B,C,D,E
}

// TeachObjDetail detail
type TeachObjDetail struct {
	ID string `json:"id"` // convert to string id

	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	TeacherID   int64     `json:"teacherId"`   // 老师id
	TeacherName string    `json:"teacherName"` // 老师名称
	CourseID    int64     `json:"courseId"`    // 课程id
	CourseName  string    `json:"courseName"`  // 课程名称
	Score       string    `json:"score"`       // 学生评价教学质量，5个等级：A,B,C,D,E
}

// CreateTeachRespond only for api docs
type CreateTeachRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		ID uint64 `json:"id"` // id
	} `json:"data"` // return data
}

// UpdateTeachByIDRespond only for api docs
type UpdateTeachByIDRespond struct {
	Result
}

// GetTeachByIDRespond only for api docs
type GetTeachByIDRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Teach TeachObjDetail `json:"teach"`
	} `json:"data"` // return data
}

// DeleteTeachByIDRespond only for api docs
type DeleteTeachByIDRespond struct {
	Result
}

// DeleteTeachsByIDsRequest request params
type DeleteTeachsByIDsRequest struct {
	IDs []uint64 `json:"ids" binding:"min=1"` // id list
}

// DeleteTeachsByIDsRespond only for api docs
type DeleteTeachsByIDsRespond struct {
	Result
}

// GetTeachByConditionRequest request params
type GetTeachByConditionRequest struct {
	query.Conditions
}

// GetTeachByConditionRespond only for api docs
type GetTeachByConditionRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Teach TeachObjDetail `json:"teach"`
	} `json:"data"` // return data
}

// ListTeachsByIDsRequest request params
type ListTeachsByIDsRequest struct {
	IDs []uint64 `json:"ids" binding:"min=1"` // id list
}

// ListTeachsByIDsRespond only for api docs
type ListTeachsByIDsRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Teachs []TeachObjDetail `json:"teachs"`
	} `json:"data"` // return data
}

// ListTeachsRequest request params
type ListTeachsRequest struct {
	query.Params
}

// ListTeachsRespond only for api docs
type ListTeachsRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Teachs []TeachObjDetail `json:"teachs"`
	} `json:"data"` // return data
}
