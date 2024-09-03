// Package initial is the package that starts the service to initialize the service, including
// the initialization configuration, service configuration, connecting to the database, and
// resource release needed when shutting down the service.
package initial

import (
	"flag"
	"strconv"

	"github.com/zhufuyi/sponge/pkg/logger"
	"github.com/zhufuyi/sponge/pkg/stat"
	"github.com/zhufuyi/sponge/pkg/tracer"

	"transfer/configs"
	"transfer/internal/config"
	"transfer/internal/rpcclient"
	//"transfer/internal/model"
)

var (
	version    string
	configFile string
)

// InitApp initial app configuration
func InitApp() {
	initConfig()
	cfg := config.Get()

	// initializing log
	_, err := logger.Init(
		logger.WithLevel(cfg.Logger.Level),
		logger.WithFormat(cfg.Logger.Format),
		logger.WithSave(cfg.Logger.IsSave),
	)
	if err != nil {
		panic(err)
	}
	logger.Debug(config.Show())
	logger.Info("[logger] was initialized")

	// initializing tracing
	if cfg.App.EnableTrace {
		tracer.InitWithConfig(
			cfg.App.Name,
			cfg.App.Env,
			cfg.App.Version,
			cfg.Jaeger.AgentHost,
			strconv.Itoa(cfg.Jaeger.AgentPort),
			cfg.App.TracingSamplingRate,
		)
		logger.Info("[tracer] was initialized")
	}

	// initializing the print system and process resources
	if cfg.App.EnableStat {
		stat.Init(
			stat.WithLog(logger.Get()),
			stat.WithAlarm(), // invalid if it is windows, the default threshold for cpu and memory is 0.8, you can modify them
		)
		logger.Info("[statistics] was initialized")
	}

	rpcclient.InitDtmServerResolver()
	rpcclient.InitTransferEndpoint()

	// initializing database
	//model.InitMysql()
	//logger.Info("init mysql succeeded")
	//model.InitCache(cfg.App.CacheType)
}

func initConfig() {
	flag.StringVar(&version, "version", "", "service Version Number")
	flag.StringVar(&configFile, "c", "", "configuration file")
	flag.Parse()

	// get configuration from local configuration file
	if configFile == "" {
		configFile = configs.Path("transfer.yml")
	}
	err := config.Init(configFile)
	if err != nil {
		panic("init config error: " + err.Error())
	}

	if version != "" {
		config.Get().App.Version = version
	}
}
