package rpcclient

import (
	"fmt"
	"strings"
	"sync"
	"time"

	_ "github.com/zhufuyi/dtmdriver-sponge"
	"google.golang.org/grpc/resolver"

	"github.com/zhufuyi/sponge/pkg/consulcli"
	"github.com/zhufuyi/sponge/pkg/etcdcli"
	"github.com/zhufuyi/sponge/pkg/logger"
	"github.com/zhufuyi/sponge/pkg/nacoscli"
	"github.com/zhufuyi/sponge/pkg/servicerd/discovery"
	"github.com/zhufuyi/sponge/pkg/servicerd/registry"
	"github.com/zhufuyi/sponge/pkg/servicerd/registry/consul"
	"github.com/zhufuyi/sponge/pkg/servicerd/registry/etcd"
	"github.com/zhufuyi/sponge/pkg/servicerd/registry/nacos"

	"stock/internal/config"
)

var (
	dtmServerEndPoint string
	dtmServerOnce     sync.Once
)

// InitDtmServerResolver init dtm resolver
func InitDtmServerResolver() {
	cfg := config.Get()

	serverName := "dtmservice"
	var grpcClientCfg config.GrpcClient
	for _, cli := range cfg.GrpcClient {
		if strings.EqualFold(cli.Name, serverName) {
			grpcClientCfg = cli
			break
		}
	}
	if grpcClientCfg.Name == "" {
		panic(fmt.Sprintf("not found grpc service name '%v' in configuration file(yaml), "+
			"please add gprc service configuration in the configuration file(yaml) under the field grpcClient", serverName))
	}

	var (
		isUseDiscover bool
		iDiscovery    registry.Discovery
	)
	switch grpcClientCfg.RegistryDiscoveryType {
	// discovering services using consul
	case "consul":
		dtmServerEndPoint = "discovery:///" + grpcClientCfg.Name // connecting to grpc services by service name
		cli, err := consulcli.Init(cfg.Consul.Addr, consulcli.WithWaitTime(time.Second*5))
		if err != nil {
			panic(fmt.Sprintf("consulcli.Init error: %v, addr: %s", err, cfg.Consul.Addr))
		}
		iDiscovery = consul.New(cli)
		isUseDiscover = true

	// discovering services using etcd
	case "etcd":
		dtmServerEndPoint = "discovery:///" + grpcClientCfg.Name // Connecting to grpc services by service name
		cli, err := etcdcli.Init(cfg.Etcd.Addrs, etcdcli.WithDialTimeout(time.Second*5))
		if err != nil {
			panic(fmt.Sprintf("etcdcli.Init error: %v, addr: %v", err, cfg.Etcd.Addrs))
		}
		iDiscovery = etcd.New(cli)
		isUseDiscover = true

	// discovering services using nacos
	case "nacos":
		// example: endpoint = "discovery:///serverName.scheme"
		dtmServerEndPoint = "discovery:///" + grpcClientCfg.Name + ".grpc" // Connecting to grpc services by service name
		cli, err := nacoscli.NewNamingClient(
			cfg.NacosRd.IPAddr,
			cfg.NacosRd.Port,
			cfg.NacosRd.NamespaceID)
		if err != nil {
			panic(fmt.Sprintf("nacoscli.NewNamingClient error: %v, ipAddr: %s, port: %d",
				err, cfg.NacosRd.IPAddr, cfg.NacosRd.Port))
		}
		iDiscovery = nacos.New(cli)
		isUseDiscover = true

	default:
		// if service discovery is not used, connect directly to the grpc service using the ip and port
		dtmServerEndPoint = fmt.Sprintf("%s:%d", grpcClientCfg.Host, grpcClientCfg.Port)
		isUseDiscover = false
	}

	if isUseDiscover {
		logger.Infof("using service discovery, type = %s, endpoint = %s", grpcClientCfg.RegistryDiscoveryType, dtmServerEndPoint)
		builder := discovery.NewBuilder(iDiscovery, discovery.WithInsecure(true), discovery.DisableDebugLog())
		resolver.Register(builder)
	} else {
		logger.Infof("using IP address connect directly, endpoint = %s", dtmServerEndPoint)
	}
}

// GetDtmEndpoint get dtm service endpoint
func GetDtmEndpoint() string {
	if dtmServerEndPoint == "" {
		dtmServerOnce.Do(func() {
			InitDtmServerResolver()
		})
	}

	return dtmServerEndPoint
}
