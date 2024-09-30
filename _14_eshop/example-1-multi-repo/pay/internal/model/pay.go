package model

import (
	"time"
)

type Pay struct {
	ID        uint64     `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	UserID    uint64     `gorm:"column:user_id;type:bigint(20) unsigned;NOT NULL" json:"userId"` // 用户id
	OrderID   string     `gorm:"column:order_id;type:varchar(45);NOT NULL" json:"orderId"`       // 订单id
	Amount    uint       `gorm:"column:amount;type:int(10) unsigned;NOT NULL" json:"amount"`     // 金额(分)
	Status    uint       `gorm:"column:status;type:int(8) unsigned;NOT NULL" json:"status"`      // 支付状态, 1:待支付, 2:已支付
	CreatedAt *time.Time `gorm:"column:created_at;type:datetime" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:datetime" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:datetime" json:"deletedAt"`
}

// TableName table name
func (m *Pay) TableName() string {
	return "pay"
}
