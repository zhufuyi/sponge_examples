package model

import (
	"time"
)

type Coupon struct {
	ID        uint64     `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	UserID    uint64     `gorm:"column:user_id;type:bigint(20) unsigned;NOT NULL" json:"userId"` // 用户id
	Amount    uint       `gorm:"column:amount;type:int(10) unsigned;NOT NULL" json:"amount"`     // 优惠券金额(分)
	Status    uint       `gorm:"column:status;type:int(8) unsigned;NOT NULL" json:"status"`      // 是否使用, 1:未使用, 2:已使用, 3:已过期
	CreatedAt *time.Time `gorm:"column:created_at;type:datetime" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:datetime" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:datetime" json:"deletedAt"`
}

// TableName table name
func (m *Coupon) TableName() string {
	return "coupon"
}
