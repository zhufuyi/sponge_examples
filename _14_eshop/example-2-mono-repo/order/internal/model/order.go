package model

import (
	"time"
)

type OrderRecord struct {
	ID           uint64     `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	CreatedAt    *time.Time `gorm:"column:created_at;type:datetime" json:"createdAt"`
	UpdatedAt    *time.Time `gorm:"column:updated_at;type:datetime" json:"updatedAt"`
	DeletedAt    *time.Time `gorm:"column:deleted_at;type:datetime" json:"deletedAt"`
	OrderID      string     `gorm:"column:order_id;type:varchar(45)" json:"orderId"`                         // 订单id
	UserID       uint64     `gorm:"column:user_id;type:bigint(20) unsigned;NOT NULL" json:"userId"`          // 用户id
	ProductID    uint64     `gorm:"column:product_id;type:bigint(20) unsigned;NOT NULL" json:"productId"`    // 商品id
	ProductCount uint       `gorm:"column:product_count;type:int(10) unsigned;NOT NULL" json:"productCount"` // 商品数量
	Amount       uint       `gorm:"column:amount;type:int(10) unsigned;NOT NULL" json:"amount"`              // 金额(分)
	CouponID     uint64     `gorm:"column:coupon_id;type:bigint(20) unsigned;NOT NULL" json:"couponId"`      // 优惠券id
	Status       uint       `gorm:"column:status;type:int(8) unsigned;NOT NULL" json:"status"`               // 订单状态: 1:待支付, 2:待发货, 3:待收货, 4:已完成, 5:已取消
}

// TableName table name
func (m *OrderRecord) TableName() string {
	return "order_record"
}
