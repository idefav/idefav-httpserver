package cfg

import (
	"encoding/json"
	"log"
	"sort"
)

const (
	OK            = "OK"
	HEALTH        = "Health"
	UNHEALTHY     = "Unhealthy"
	VERSION       = "VERSION"
	ERROR_HANDLER = "ERROR"
)

type ServerConfig struct {
	// app name
	AppName string
	// listen Address
	Address string
	// 优雅停机
	GracefulShutdown bool
	// 优雅停机等待时间(毫秒)
	GracefulShutdownWaitTimeMs int
	// 预热启动
	Warmup bool
	// 路由组件名称
	RouterName string
	// 访问日志
	AccessLog bool
	// 配置文件
	ConfigFile string
}

type ServerConfigLoader interface {
	GetName() string
	GetOrder() int
	Load(config *ServerConfig)
}

type ServerConfigLoaders []ServerConfigLoader

func (s ServerConfigLoaders) Len() int {
	return len(s)
}

func (s ServerConfigLoaders) Less(i, j int) bool {
	return s[j].GetOrder() > s[i].GetOrder()
}

func (s ServerConfigLoaders) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

var ServerConfigLoaderList = ServerConfigLoaders{}

func SetUp() *ServerConfig {
	config := ServerConfig{
		AppName:                    "idefav-httpserver",
		Address:                    ":8080",
		GracefulShutdown:           true,
		GracefulShutdownWaitTimeMs: 30 * 1000,
	}
	// 加载配置
	loadConfig(&config)
	return &config
}

// 加载配置, 按照 Order 从小到大排序, Order 越大优先级越高
func loadConfig(config *ServerConfig) {
	sort.Sort(ServerConfigLoaderList)
	for _, loader := range ServerConfigLoaderList {
		log.Println("load config from " + loader.GetName())
		loader.Load(config)
	}
	printServerConfig(config)
}

func printServerConfig(config *ServerConfig) {
	data, _ := json.MarshalIndent(config, "", "  ")
	log.Printf("Loaded Config: \n%v", string(data))
}

func AddServerConfigLoader(configLoader ServerConfigLoader) {
	ServerConfigLoaderList = append(ServerConfigLoaderList, configLoader)
}
