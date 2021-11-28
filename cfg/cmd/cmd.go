package cmd

import (
	"flag"
	"idefav-httpserver/cfg"
	"math"
	"strconv"
)

const (
	ADDRESS                        = "address"
	SHUTDOWN_GRACEFUL              = "shutdown.graceful"
	SHUTDOWN_GRACEFUL_WAIT_TIME_MS = "shutdown.graceful.wait-time-ms"
	WARMUP                         = "warmup"
	ROUTER_NAME                    = "router.name"
	ACCESS_LOG                     = "access.log"
	CONFIG_FILE                    = "config.file"
)

type Command struct {
	Name  string
	Order int
}

func (c Command) GetName() string {
	return c.Name
}

func (c Command) GetOrder() int {
	return c.Order
}

func (c Command) Load(config *cfg.ServerConfig) {
	address := flag.String(ADDRESS, "", "设置服务监听地址")
	shutdownGraceful := flag.String(SHUTDOWN_GRACEFUL, "", "是否优雅停机")
	shutdownGracefulWaitTimeMs := flag.String(SHUTDOWN_GRACEFUL_WAIT_TIME_MS, "", "设置优雅停机等待时间")
	warmup := flag.String(WARMUP, "", "启动预热")
	routerName := flag.String(ROUTER_NAME, "", "路由组件")
	accessLog := flag.String(ACCESS_LOG, "", "开启访问日志")
	flag.Parse()
	if *address != "" {
		config.Address = *address
	}
	if *shutdownGraceful != "" {
		if val, err := strconv.ParseBool(*shutdownGraceful); err == nil {
			config.GracefulShutdown = val
		}
	}
	if *shutdownGracefulWaitTimeMs != "" {
		if val, err := strconv.ParseInt(*shutdownGracefulWaitTimeMs, 10, 0); err == nil {
			config.GracefulShutdownWaitTimeMs = int(val)
		}
	}

	if *warmup != "" {
		if val, err := strconv.ParseBool(*warmup); err == nil {
			config.Warmup = val
		}
	}
	if *routerName != "" {
		config.RouterName = *routerName
	}

	if *accessLog != "" {
		if val, err := strconv.ParseBool(*accessLog); err == nil {
			config.AccessLog = val
		}
	}

}

func init() {
	cmd := &Command{
		Name:  "command",
		Order: math.MaxInt,
	}
	cfg.AddServerConfigLoader(cmd)
}
