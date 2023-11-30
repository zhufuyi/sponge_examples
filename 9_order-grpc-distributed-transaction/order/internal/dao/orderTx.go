package dao

import (
	"database/sql"

	"github.com/zhufuyi/sponge/pkg/logger"

	"github.com/dtm-labs/client/dtmcli/dtmimp"
)

// CreateOrder 创建订单
func CreateOrder(tx *sql.Tx, orderId string, userId uint64, productId uint64, amount uint32, productCount uint32, couponId uint64) error {
	sqlStr := "insert into eshop_order.order_record(order_id, user_id, product_id, amount, product_count, coupon_id, status, created_at, updated_at) values(?,?,?,?,?,?,1,now(),now())" // 1: 待支付
	sqlFieldValues := []interface{}{orderId, userId, productId, amount, productCount, couponId}
	logger.Info("sql", logger.String("sql", sqlStr), logger.Any("fieldValues", sqlFieldValues))
	_, err := dtmimp.DBExec(dtmimp.DBTypeMysql, tx, sqlStr, sqlFieldValues...)
	return err
}

// CreateOrderRevert 取消创建订单
func CreateOrderRevert(tx *sql.Tx, orderId string) error {
	sqlStr := "update eshop_order.order_record set status=5, updated_at=now() where order_id = ?" // 5: 取消创建订单
	sqlFieldValues := []interface{}{orderId}
	logger.Info("sql", logger.String("sql", sqlStr), logger.Any("fieldValues", sqlFieldValues))
	_, err := dtmimp.DBExec(dtmimp.DBTypeMysql, tx, sqlStr, sqlFieldValues...)
	return err
}
