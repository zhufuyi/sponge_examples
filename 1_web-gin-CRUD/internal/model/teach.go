package model

import (
	"github.com/zhufuyi/sponge/pkg/mysql"
)

// Teach 老师课程
type Teach struct {
	mysql.Model `gorm:"embedded"`

	TeacherID   int64  `gorm:"column:teacher_id;type:bigint(20);NOT NULL" json:"teacherId"`      // 老师id
	TeacherName string `gorm:"column:teacher_name;type:varchar(50);NOT NULL" json:"teacherName"` // 老师名称
	CourseID    int64  `gorm:"column:course_id;type:bigint(20);NOT NULL" json:"courseId"`        // 课程id
	CourseName  string `gorm:"column:course_name;type:varchar(50);NOT NULL" json:"courseName"`   // 课程名称
	Score       string `gorm:"column:score;type:char(5);NOT NULL" json:"score"`                  // 学生评价教学质量，5个等级：A,B,C,D,E
}

// TableName table name
func (m *Teach) TableName() string {
	return "teach"
}
