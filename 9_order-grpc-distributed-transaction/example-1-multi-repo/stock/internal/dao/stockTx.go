package dao

import (
	"database/sql"

	"github.com/zhufuyi/sponge/pkg/logger"

	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmcli/dtmimp"
)

// StockDeduct 扣减库存
func StockDeduct(tx *sql.Tx, productID uint64, productCount uint32) error {
	sqlStr := "update eshop_stock.stock set stock=stock-?, updated_at=now() where product_id = ? and stock >= ?"
	sqlFieldValues := []interface{}{productCount, productID, productCount}
	logger.Info("sql", logger.String("sql", sqlStr), logger.Any("fieldValues", sqlFieldValues))
	affected, err := dtmimp.DBExec(dtmimp.DBTypeMysql, tx, sqlStr, sqlFieldValues...)
	if err == nil && affected == 0 {
		return dtmcli.ErrFailure // global transactions will be rolled back
	}
	return err
}

// StockDeductRevert 补偿库存
func StockDeductRevert(tx *sql.Tx, productID uint64, productCount uint32) error {
	sqlStr := "update eshop_stock.stock set stock=stock+?, updated_at=now() where product_id = ?"
	sqlFieldValues := []interface{}{productCount, productID}
	logger.Info("sql", logger.String("sql", sqlStr), logger.Any("fieldValues", sqlFieldValues))
	_, err := dtmimp.DBExec(dtmimp.DBTypeMysql, tx, sqlStr, sqlFieldValues...)
	return err
}
