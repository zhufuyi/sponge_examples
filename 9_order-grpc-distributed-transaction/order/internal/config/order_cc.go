// code generated by https://order

package config

import (
	"github.com/zhufuyi/sponge/pkg/conf"
)

func NewCenter(configFile string) (*Center, error) {
	nacosConf := &Center{}
	err := conf.Parse(configFile, nacosConf)
	return nacosConf, err
}

type Center struct {
	Nacos Nacos `yaml:"nacos" json:"nacos"`
}

type Nacos struct {
	ContextPath string `yaml:"contextPath" json:"contextPath"`
	DataID      string `yaml:"dataID" json:"dataID"`
	Format      string `yaml:"format" json:"format"`
	Group       string `yaml:"group" json:"group"`
	IPAddr      string `yaml:"ipAddr" json:"ipAddr"`
	NamespaceID string `yaml:"namespaceID" json:"namespaceID"`
	Port        int    `yaml:"port" json:"port"`
	Scheme      string `yaml:"scheme" json:"scheme"`
}
