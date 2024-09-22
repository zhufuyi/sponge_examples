package model

import (
	"time"
)

type Stock struct {
	ID        uint64     `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	ProductID uint64     `gorm:"column:product_id;type:bigint(20) unsigned;NOT NULL" json:"productId"` // 商品id
	Stock     uint       `gorm:"column:stock;type:int(11) unsigned;NOT NULL" json:"stock"`             // 库存
	CreatedAt *time.Time `gorm:"column:created_at;type:datetime" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:datetime" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:datetime" json:"deletedAt"`
}

// TableName table name
func (m *Stock) TableName() string {
	return "stock"
}
