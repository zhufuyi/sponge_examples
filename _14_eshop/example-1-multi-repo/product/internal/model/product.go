package model

import (
	"github.com/zhufuyi/sponge/pkg/ggorm"
)

type Product struct {
	ggorm.Model `gorm:"embedded"` // embed id and time

	Name        string `gorm:"column:name;type:varchar(255)" json:"name"`       // 商品名称
	Price       int    `gorm:"column:price;type:int(11)" json:"price"`          // 商品价格
	Photo       string `gorm:"column:photo;type:varchar(255)" json:"photo"`     // 商品图片
	Description string `gorm:"column:description;type:text" json:"description"` // 商品描述
}

// TableName table name
func (m *Product) TableName() string {
	return "product"
}
