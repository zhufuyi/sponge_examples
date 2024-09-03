package rpcclient

import (
	"fmt"
	"strings"
	"sync"

	"github.com/zhufuyi/sponge/pkg/logger"

	"flashSale/internal/config"
)

var (
	flashSaleEndPoint string
	flashSaleOnce     sync.Once
)

// get endpoint for dtm service
func getEndpoint(serverName string) (string, error) {
	cfg := config.Get()

	var grpcClientCfg config.GrpcClient

	// local service
	if serverName == cfg.App.Name {
		grpcClientCfg = config.GrpcClient{
			Name:                  serverName,
			Host:                  cfg.App.Host,
			Port:                  cfg.Grpc.Port,
			RegistryDiscoveryType: cfg.App.RegistryDiscoveryType,
			EnableLoadBalance:     true,
		}
	} else {
		// remote service
		for _, cli := range cfg.GrpcClient {
			if strings.EqualFold(cli.Name, serverName) {
				grpcClientCfg = cli
				break
			}
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
		logger.Infof("[dtm] connects directly to the [%s] service through IP address, endpoint = %s", serverName, endpoint)
	}

	return endpoint, nil
}

// InitEndpointsForDtm init endpoints for dtm service
func InitEndpointsForDtm() {
	GetFlashSaleEndpoint()
}

// GetFlashSaleEndpoint get flash sale grpc endpoint
func GetFlashSaleEndpoint() string {
	if flashSaleEndPoint == "" {
		flashSaleOnce.Do(func() {
			endpoint, err := getEndpoint(config.Get().App.Name)
			if err != nil {
				panic(err)
			}
			flashSaleEndPoint = endpoint
		})
	}

	return flashSaleEndPoint
}
