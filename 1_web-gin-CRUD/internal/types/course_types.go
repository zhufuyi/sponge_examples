// Package types define the structure of request parameters and respond results in this package
package types

import (
	"time"

	"github.com/zhufuyi/sponge/pkg/mysql/query"
)

var _ time.Time

// Tip: suggested filling in the binding rules https://github.com/go-playground/validator in request struct fields tag.

// CreateCourseRequest request params
type CreateCourseRequest struct {
	Code     string `json:"code" binding:""`     // 课程代号
	Name     string `json:"name" binding:""`     // 课程名称
	Credit   int    `json:"credit" binding:""`   // 学分
	College  string `json:"college" binding:""`  // 学院
	Semester string `json:"semester" binding:""` // 学期
	Time     string `json:"time" binding:""`     // 上课时间
	Place    string `json:"place" binding:""`    // 上课地点
}

// UpdateCourseByIDRequest request params
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

// CourseObjDetail detail
type CourseObjDetail struct {
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

// CreateCourseRespond only for api docs
type CreateCourseRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		ID uint64 `json:"id"` // id
	} `json:"data"` // return data
}

// UpdateCourseByIDRespond only for api docs
type UpdateCourseByIDRespond struct {
	Result
}

// GetCourseByIDRespond only for api docs
type GetCourseByIDRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Course CourseObjDetail `json:"course"`
	} `json:"data"` // return data
}

// DeleteCourseByIDRespond only for api docs
type DeleteCourseByIDRespond struct {
	Result
}

// DeleteCoursesByIDsRequest request params
type DeleteCoursesByIDsRequest struct {
	IDs []uint64 `json:"ids" binding:"min=1"` // id list
}

// DeleteCoursesByIDsRespond only for api docs
type DeleteCoursesByIDsRespond struct {
	Result
}

// GetCourseByConditionRequest request params
type GetCourseByConditionRequest struct {
	query.Conditions
}

// GetCourseByConditionRespond only for api docs
type GetCourseByConditionRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Course CourseObjDetail `json:"course"`
	} `json:"data"` // return data
}

// ListCoursesByIDsRequest request params
type ListCoursesByIDsRequest struct {
	IDs []uint64 `json:"ids" binding:"min=1"` // id list
}

// ListCoursesByIDsRespond only for api docs
type ListCoursesByIDsRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Courses []CourseObjDetail `json:"courses"`
	} `json:"data"` // return data
}

// ListCoursesRequest request params
type ListCoursesRequest struct {
	query.Params
}

// ListCoursesRespond only for api docs
type ListCoursesRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Courses []CourseObjDetail `json:"courses"`
	} `json:"data"` // return data
}
