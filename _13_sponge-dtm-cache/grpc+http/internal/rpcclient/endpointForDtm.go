package rpcclient

import (
	"fmt"
	"strings"
	"sync"

	"github.com/zhufuyi/sponge/pkg/logger"

	"stock/internal/config"
)

var (
	stockEndPoint string
	stockOnce     sync.Once
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
	GetStockEndpoint()
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
