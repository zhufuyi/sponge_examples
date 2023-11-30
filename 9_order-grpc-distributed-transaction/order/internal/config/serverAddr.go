package config

import (
	"github.com/zhufuyi/sponge/pkg/utils"
)

var DtmGrpcAddr, CouponGrpcAddr, StockGrpcAddr, PayGrpcAddr, OrderGrpcAddr string

// SetServerAddr 从配置中获取服务地址
func SetServerAddr() {
	DtmGrpcAddr = Get().Dtm.Addr

	for _, client := range Get().GrpcClient {
		switch client.Name {
		case "order":
			OrderGrpcAddr = client.Host + ":" + utils.IntToStr(client.Port)
		case "stock":
			StockGrpcAddr = client.Host + ":" + utils.IntToStr(client.Port)
		case "coupon":
			CouponGrpcAddr = client.Host + ":" + utils.IntToStr(client.Port)
		case "pay":
			PayGrpcAddr = client.Host + ":" + utils.IntToStr(client.Port)
		}
	}
}
