package dao

import (
	"database/sql"

	"github.com/zhufuyi/sponge/pkg/logger"

	"github.com/dtm-labs/client/dtmcli/dtmimp"
)

// CreatePay 创建支付订单
func CreatePay(tx *sql.Tx, userId uint64, orderId string, amount uint32) error {
	sqlStr := "insert into eshop_pay.pay(user_id, order_id, amount, status, created_at, updated_at) values(?,?,?,1,now(),now())" // 1: 待支付
	sqlFieldValues := []interface{}{userId, orderId, amount}
	logger.Info("sql", logger.String("sql", sqlStr), logger.Any("fieldValues", sqlFieldValues))
	_, err := dtmimp.DBExec(dtmimp.DBTypeMysql, tx, sqlStr, sqlFieldValues...)
	return err
}

// CreatePayRevert 取消支付订单
func CreatePayRevert(tx *sql.Tx, orderId string) error {
	sqlStr := "update eshop_pay.pay set status=4, updated_at=now() where order_id = ?" // 4: 取消支付订单
	sqlFieldValues := []interface{}{orderId}
	logger.Info("sql", logger.String("sql", sqlStr), logger.Any("fieldValues", sqlFieldValues))
	_, err := dtmimp.DBExec(dtmimp.DBTypeMysql, tx, sqlStr, sqlFieldValues...)
	return err
}
