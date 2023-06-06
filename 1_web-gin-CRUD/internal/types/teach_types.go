package types

import (
	"time"

	"github.com/zhufuyi/sponge/pkg/mysql/query"
)

var _ time.Time

// CreateTeachRequest create params
// todo fill in the binding rules https://github.com/go-playground/validator
type CreateTeachRequest struct {
	TeacherID   int64  `json:"teacherId" binding:""`   // 老师id
	TeacherName string `json:"teacherName" binding:""` // 老师名称
	CourseID    int64  `json:"courseId" binding:""`    // 课程id
	CourseName  string `json:"courseName" binding:""`  // 课程名称
	Score       string `json:"score" binding:""`       // 学生评价教学质量，5个等级：A,B,C,D,E
}

// UpdateTeachByIDRequest update params
type UpdateTeachByIDRequest struct {
	ID uint64 `json:"id" binding:""` // uint64 id

	TeacherID   int64  `json:"teacherId" binding:""`   // 老师id
	TeacherName string `json:"teacherName" binding:""` // 老师名称
	CourseID    int64  `json:"courseId" binding:""`    // 课程id
	CourseName  string `json:"courseName" binding:""`  // 课程名称
	Score       string `json:"score" binding:""`       // 学生评价教学质量，5个等级：A,B,C,D,E
}

// GetTeachByIDRespond respond detail
type GetTeachByIDRespond struct {
	ID string `json:"id"` // convert to string id

	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	TeacherID   int64     `json:"teacherId"`   // 老师id
	TeacherName string    `json:"teacherName"` // 老师名称
	CourseID    int64     `json:"courseId"`    // 课程id
	CourseName  string    `json:"courseName"`  // 课程名称
	Score       string    `json:"score"`       // 学生评价教学质量，5个等级：A,B,C,D,E
}

// DeleteTeachsByIDsRequest request form ids
type DeleteTeachsByIDsRequest struct {
	IDs []uint64 `json:"ids" binding:"min=1"` // id list
}

// GetTeachsByIDsRequest request form ids
type GetTeachsByIDsRequest struct {
	IDs []uint64 `json:"ids" binding:"min=1"` // id list
}

// GetTeachsRequest request form params
type GetTeachsRequest struct {
	query.Params // query parameters
}

// ListTeachsRespond list data
type ListTeachsRespond []struct {
	GetTeachByIDRespond
}
