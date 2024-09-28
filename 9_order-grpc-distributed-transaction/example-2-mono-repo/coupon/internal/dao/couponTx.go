package dao

import (
	"database/sql"

	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmcli/dtmimp"

	"github.com/zhufuyi/sponge/pkg/logger"
)

// CouponUse 使用消费券
func CouponUse(tx *sql.Tx, id uint64) error {
	sqlStr := "update eshop_coupon.coupon set status=2, updated_at=now() where id=? and status=1" // 1:未使用, 2:已使用
	sqlFieldValues := []interface{}{id}
	logger.Info("sql", logger.String("sql", sqlStr), logger.Any("fieldValues", sqlFieldValues))
	affected, err := dtmimp.DBExec(dtmimp.DBTypeMysql, tx, sqlStr, sqlFieldValues...)
	if err == nil && affected == 0 {
		return dtmcli.ErrFailure // global transactions will be rolled back
	}
	return err
}

// CouponUseRevert 补偿消费券
func CouponUseRevert(tx *sql.Tx, id uint64) error {
	sqlStr := "update eshop_coupon.coupon set status=1, updated_at=now() where id=? and status=2" // 1:未使用, 2:已使用
	sqlFieldValues := []interface{}{id}
	logger.Info("sql", logger.String("sql", sqlStr), logger.Any("fieldValues", sqlFieldValues))
	_, err := dtmimp.DBExec(dtmimp.DBTypeMysql, tx, sqlStr, sqlFieldValues...)
	return err
}
