package rpcclient

import (
	"fmt"
	"sync"

	"transfer/internal/config"
)

var (
	transferEndPoint string
	transferOnce     sync.Once
)

// InitTransferEndpoint init transfer endpoint
func InitTransferEndpoint() {
	switch config.Get().App.RegistryDiscoveryType {
	case "consul", "etcd":
		transferEndPoint = "discovery:///" + config.Get().App.Name
	case "nacos":
		transferEndPoint = "discovery:///" + config.Get().App.Name + ".grpc"
	default:
		transferEndPoint = fmt.Sprintf("%s:%d", config.Get().App.Host, config.Get().Grpc.Port)
	}
}

// GetTransferEndpoint get transfer endpoint
func GetTransferEndpoint() string {
	if transferEndPoint == "" {
		transferOnce.Do(func() {
			InitTransferEndpoint()
		})
	}

	return transferEndPoint
}
