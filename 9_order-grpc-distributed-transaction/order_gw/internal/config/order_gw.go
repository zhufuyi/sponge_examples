// Package config is a package that parses configuration files into structures, and supports
// parsing yaml, json and toml files., code generated by https://order_gw.
package config

import (
	"github.com/zhufuyi/sponge/pkg/conf"
)

var config *Config

func Init(configFile string, fs ...func()) error {
	config = &Config{}
	return conf.Parse(configFile, config, fs...)
}

func Show(hiddenFields ...string) string {
	return conf.Show(config, hiddenFields...)
}

func Get() *Config {
	if config == nil {
		panic("config is nil")
	}
	return config
}

func Set(conf *Config) {
	config = conf
}

type Config struct {
	App        App          `yaml:"app" json:"app"`
	Consul     Consul       `yaml:"consul" json:"consul"`
	Etcd       Etcd         `yaml:"etcd" json:"etcd"`
	Grpc       Grpc         `yaml:"grpc" json:"grpc"`
	GrpcClient []GrpcClient `yaml:"grpcClient" json:"grpcClient"`
	HTTP       HTTP         `yaml:"http" json:"http"`
	Jaeger     Jaeger       `yaml:"jaeger" json:"jaeger"`
	Logger     Logger       `yaml:"logger" json:"logger"`
	Mysql      Mysql        `yaml:"mysql" json:"mysql"`
	NacosRd    NacosRd      `yaml:"nacosRd" json:"nacosRd"`
	Redis      Redis        `yaml:"redis" json:"redis"`
}

type Consul struct {
	Addr string `yaml:"addr" json:"addr"`
}

type Etcd struct {
	Addrs []string `yaml:"addrs" json:"addrs"`
}

type Jaeger struct {
	AgentHost string `yaml:"agentHost" json:"agentHost"`
	AgentPort int    `yaml:"agentPort" json:"agentPort"`
}

type ClientToken struct {
	AppID  string `yaml:"appID" json:"appID"`
	AppKey string `yaml:"appKey" json:"appKey"`
	Enable bool   `yaml:"enable" json:"enable"`
}

type ClientSecure struct {
	CaFile     string `yaml:"caFile" json:"caFile"`
	CertFile   string `yaml:"certFile" json:"certFile"`
	KeyFile    string `yaml:"keyFile" json:"keyFile"`
	ServerName string `yaml:"serverName" json:"serverName"`
	Type       string `yaml:"type" json:"type"`
}

type ServerSecure struct {
	CaFile   string `yaml:"caFile" json:"caFile"`
	CertFile string `yaml:"certFile" json:"certFile"`
	KeyFile  string `yaml:"keyFile" json:"keyFile"`
	Type     string `yaml:"type" json:"type"`
}

type App struct {
	CacheType             string  `yaml:"cacheType" json:"cacheType"`
	EnableCircuitBreaker  bool    `yaml:"enableCircuitBreaker" json:"enableCircuitBreaker"`
	EnableHTTPProfile     bool    `yaml:"enableHTTPProfile" json:"enableHTTPProfile"`
	EnableLimit           bool    `yaml:"enableLimit" json:"enableLimit"`
	EnableMetrics         bool    `yaml:"enableMetrics" json:"enableMetrics"`
	EnableStat            bool    `yaml:"enableStat" json:"enableStat"`
	EnableTrace           bool    `yaml:"enableTrace" json:"enableTrace"`
	Env                   string  `yaml:"env" json:"env"`
	Host                  string  `yaml:"host" json:"host"`
	Name                  string  `yaml:"name" json:"name"`
	RegistryDiscoveryType string  `yaml:"registryDiscoveryType" json:"registryDiscoveryType"`
	TracingSamplingRate   float64 `yaml:"tracingSamplingRate" json:"tracingSamplingRate"`
	Version               string  `yaml:"version" json:"version"`
}

type GrpcClient struct {
	ClientSecure          ClientSecure `yaml:"clientSecure" json:"clientSecure"`
	ClientToken           ClientToken  `yaml:"clientToken" json:"clientToken"`
	EnableLoadBalance     bool         `yaml:"enableLoadBalance" json:"enableLoadBalance"`
	Host                  string       `yaml:"host" json:"host"`
	Name                  string       `yaml:"name" json:"name"`
	Port                  int          `yaml:"port" json:"port"`
	RegistryDiscoveryType string       `yaml:"registryDiscoveryType" json:"registryDiscoveryType"`
}

type Mysql struct {
	ConnMaxLifetime int      `yaml:"connMaxLifetime" json:"connMaxLifetime"`
	Dsn             string   `yaml:"dsn" json:"dsn"`
	EnableLog       bool     `yaml:"enableLog" json:"enableLog"`
	MastersDsn      []string `yaml:"mastersDsn" json:"mastersDsn"`
	MaxIdleConns    int      `yaml:"maxIdleConns" json:"maxIdleConns"`
	MaxOpenConns    int      `yaml:"maxOpenConns" json:"maxOpenConns"`
	SlavesDsn       []string `yaml:"slavesDsn" json:"slavesDsn"`
	SlowThreshold   int      `yaml:"slowThreshold" json:"slowThreshold"`
}

type Redis struct {
	DialTimeout  int    `yaml:"dialTimeout" json:"dialTimeout"`
	Dsn          string `yaml:"dsn" json:"dsn"`
	ReadTimeout  int    `yaml:"readTimeout" json:"readTimeout"`
	WriteTimeout int    `yaml:"writeTimeout" json:"writeTimeout"`
}

type Grpc struct {
	EnableToken  bool         `yaml:"enableToken" json:"enableToken"`
	HTTPPort     int          `yaml:"httpPort" json:"httpPort"`
	Port         int          `yaml:"port" json:"port"`
	ReadTimeout  int          `yaml:"readTimeout" json:"readTimeout"`
	ServerSecure ServerSecure `yaml:"serverSecure" json:"serverSecure"`
	WriteTimeout int          `yaml:"writeTimeout" json:"writeTimeout"`
}

type Logger struct {
	Format string `yaml:"format" json:"format"`
	IsSave bool   `yaml:"isSave" json:"isSave"`
	Level  string `yaml:"level" json:"level"`
}

type NacosRd struct {
	IPAddr      string `yaml:"ipAddr" json:"ipAddr"`
	NamespaceID string `yaml:"namespaceID" json:"namespaceID"`
	Port        int    `yaml:"port" json:"port"`
}

type HTTP struct {
	Port         int `yaml:"port" json:"port"`
	ReadTimeout  int `yaml:"readTimeout" json:"readTimeout"`
	WriteTimeout int `yaml:"writeTimeout" json:"writeTimeout"`
}
