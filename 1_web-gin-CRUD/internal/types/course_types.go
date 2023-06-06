package types

import (
	"time"

	"github.com/zhufuyi/sponge/pkg/mysql/query"
)

var _ time.Time

// CreateCourseRequest create params
// todo fill in the binding rules https://github.com/go-playground/validator
type CreateCourseRequest struct {
	Code     string `json:"code" binding:""`     // 课程代号
	Name     string `json:"name" binding:""`     // 课程名称
	Credit   int    `json:"credit" binding:""`   // 学分
	College  string `json:"college" binding:""`  // 学院
	Semester string `json:"semester" binding:""` // 学期
	Time     string `json:"time" binding:""`     // 上课时间
	Place    string `json:"place" binding:""`    // 上课地点
}

// UpdateCourseByIDRequest update params
type UpdateCourseByIDRequest struct {
	ID uint64 `json:"id" binding:""` // uint64 id

	Code     string `json:"code" binding:""`     // 课程代号
	Name     string `json:"name" binding:""`     // 课程名称
	Credit   int    `json:"credit" binding:""`   // 学分
	College  string `json:"college" binding:""`  // 学院
	Semester string `json:"semester" binding:""` // 学期
	Time     string `json:"time" binding:""`     // 上课时间
	Place    string `json:"place" binding:""`    // 上课地点
}

// GetCourseByIDRespond respond detail
type GetCourseByIDRespond struct {
	ID string `json:"id"` // convert to string id

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Code      string    `json:"code"`     // 课程代号
	Name      string    `json:"name"`     // 课程名称
	Credit    int       `json:"credit"`   // 学分
	College   string    `json:"college"`  // 学院
	Semester  string    `json:"semester"` // 学期
	Time      string    `json:"time"`     // 上课时间
	Place     string    `json:"place"`    // 上课地点
}

// DeleteCoursesByIDsRequest request form ids
type DeleteCoursesByIDsRequest struct {
	IDs []uint64 `json:"ids" binding:"min=1"` // id list
}

// GetCoursesByIDsRequest request form ids
type GetCoursesByIDsRequest struct {
	IDs []uint64 `json:"ids" binding:"min=1"` // id list
}

// GetCoursesRequest request form params
type GetCoursesRequest struct {
	query.Params // query parameters
}

// ListCoursesRespond list data
type ListCoursesRespond []struct {
	GetCourseByIDRespond
}
