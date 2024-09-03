package rpcclient

import (
	"fmt"
	"strings"
	"sync"

	"github.com/zhufuyi/sponge/pkg/logger"

	"transfer/internal/config"
)

var (
	transferEndPoint string
	transferOnce     sync.Once
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

// InitTransferEndpoint init transfer endpoint
func InitTransferEndpoint() {
	GetTransferEndpoint()
}

// GetTransferEndpoint get transfer endpoint
func GetTransferEndpoint() string {
	if transferEndPoint == "" {
		transferOnce.Do(func() {
			endpoint, err := getEndpoint("transfer")
			if err != nil {
				panic(err)
			}
			transferEndPoint = endpoint
		})
	}

	return transferEndPoint
}
