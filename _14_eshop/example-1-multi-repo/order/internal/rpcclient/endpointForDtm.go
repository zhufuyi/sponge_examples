package rpcclient

import (
	"fmt"
	"strings"
	"sync"

	"github.com/zhufuyi/sponge/pkg/logger"

	"order/internal/config"
)

var (
	orderEndPoint string
	orderOnce     sync.Once

	couponEndPoint string
	couponOnce2    sync.Once

	stockEndPoint string
	stockOnce     sync.Once

	payEndPoint string
	payOnce     sync.Once
)

// get endpoint for dtm service
func getEndpoint(serverName string) (string, error) {
	cfg := config.Get()

	var grpcClientCfg config.GrpcClient
	for _, cli := range cfg.GrpcClient {
		if strings.EqualFold(cli.Name, serverName) {
			grpcClientCfg = cli
			break
		}
	}
	if grpcClientCfg.Name == "" {
		return "", fmt.Errorf("not found grpc service name '%v' in configuration file(yaml), "+
			"please add gprc service configuration in the configuration file(yaml) under the field grpcClient", serverName)
	}

	var endpoint string
	var isUseDiscover bool
	switch grpcClientCfg.RegistryDiscoveryType {
	case "consul", "etcd":
		endpoint = "discovery:///" + grpcClientCfg.Name
		isUseDiscover = true
	case "nacos":
		endpoint = "discovery:///" + grpcClientCfg.Name + ".grpc"
		isUseDiscover = true
	default:
		endpoint = fmt.Sprintf("%s:%d", grpcClientCfg.Host, grpcClientCfg.Port)
	}

	if isUseDiscover {
		logger.Infof("[dtm] connects to the [%s] service through service discovery, type = %s, endpoint = %s", serverName, grpcClientCfg.RegistryDiscoveryType, endpoint)
	} else {
		logger.Infof("[dtm] connects to the [%s] service through IP address, endpoint = %s", serverName, endpoint)
	}

	return endpoint, nil
}

// InitEndpointsForDtm init endpoints for dtm service
func InitEndpointsForDtm() {
	GetOrderEndpoint()
	GetCouponEndpoint()
	GetStockEndpoint()
	GetPayEndpoint()
}

// GetOrderEndpoint get order endpoint
func GetOrderEndpoint() string {
	if orderEndPoint == "" {
		orderOnce.Do(func() {
			endpoint, err := getEndpoint("order")
			if err != nil {
				panic(err)
			}
			orderEndPoint = endpoint
		})
	}

	return orderEndPoint
}

// GetCouponEndpoint get coupon endpoint
func GetCouponEndpoint() string {
	if couponEndPoint == "" {
		couponOnce2.Do(func() {
			endpoint, err := getEndpoint("coupon")
			if err != nil {
				panic(err)
			}
			couponEndPoint = endpoint
		})
	}

	return couponEndPoint
}

// GetStockEndpoint get stock endpoint
func GetStockEndpoint() string {
	if stockEndPoint == "" {
		stockOnce.Do(func() {
			endpoint, err := getEndpoint("stock")
			if err != nil {
				panic(err)
			}
			stockEndPoint = endpoint
		})
	}

	return stockEndPoint
}

// GetPayEndpoint get pay endpoint
func GetPayEndpoint() string {
	if payEndPoint == "" {
		payOnce.Do(func() {
			endpoint, err := getEndpoint("pay")
			if err != nil {
				panic(err)
			}
			payEndPoint = endpoint
		})
	}

	return payEndPoint
}
