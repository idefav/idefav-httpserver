package yaml

import (
	"encoding/json"
	"flag"
	"idefav-httpserver/cfg"
	"idefav-httpserver/cfg/cmd"
	"idefav-httpserver/cfg/env"
	"log"
	"math"
	"os"
)
import yml "gopkg.in/yaml.v2"

//type YamlConfig struct {
//	// listen Address
//	Address string `yaml:"server.address"`
//	// 优雅停机
//	GracefulShutdown bool `yaml:"server.shutdown.graceful.enabled"`
//	// 优雅停机等待时间(毫秒)
//	GracefulShutdownWaitTimeMs int `yaml:"server.shutdown.graceful.wait-time-ms"`
//	// 预热启动
//	Warmup bool `yaml:"server.warmup"`
//	// 路由组件名称
//	RouterName string `yaml:"router.name"`
//	// 访问日志
//	AccessLog bool `yaml:"server.access-log.enabled"`
//}

type YamlConfig struct {
	Server struct {
		Address   string `yaml:"address"`
		Warmup    bool   `yaml:"warmup"`
		AccessLog struct {
			Enabled bool `yaml:"enabled"`
		} `yaml:"access-log"`
		Shutdown struct {
			Graceful struct {
				Enabled    bool `yaml:"enabled"`
				WaitTimeMs int  `yaml:"wait-time-ms"`
			} `yaml:"graceful"`
		} `yaml:"shutdown"`
	} `yaml:"server"`
	Router struct {
		Name string `yaml:"name"`
	} `yaml:"router"`
}

func (y YamlConfig) GetName() string {
	return "Yaml"
}

func (y YamlConfig) GetOrder() int {
	return math.MaxInt - 2
}

func (y YamlConfig) Load(config *cfg.ServerConfig) {
	// 加载配置文件路径
	if config.ConfigFile == "" {
		c := flag.String(cmd.CONFIG_FILE, "server.yaml", "配置文件路径")
		flag.Parse()
		config.ConfigFile = *c
	}

	if config.ConfigFile == "" {
		if val := env.GetStringEnvVal(env.CONFIG_FILE); val != "" {
			config.ConfigFile = val
		}
	}

	if config.ConfigFile == "" {
		config.ConfigFile = "server.yaml"
	}
	yamlConfig, err := ReadYamlConfig(config.ConfigFile)
	if err != nil {
		log.Println("Load yaml config file failed, ignore")
		return
	}
	marshal, _ := json.MarshalIndent(yamlConfig, "", "  ")
	log.Printf("yaml: %s\n", string(marshal))
	config.Address = yamlConfig.Server.Address
	config.GracefulShutdown = yamlConfig.Server.Shutdown.Graceful.Enabled
	config.RouterName = yamlConfig.Router.Name
	config.Warmup = yamlConfig.Server.Warmup
	config.AccessLog = yamlConfig.Server.AccessLog.Enabled
	config.GracefulShutdownWaitTimeMs = yamlConfig.Server.Shutdown.Graceful.WaitTimeMs
}

func ReadYamlConfig(path string) (*YamlConfig, error) {
	conf := &YamlConfig{}
	if f, err := os.Open(path); err != nil {
		return nil, err
	} else {
		err2 := yml.NewDecoder(f).Decode(conf)
		if err2 != nil {
			return nil, err2
		}
	}
	return conf, nil
}

func init() {
	config := YamlConfig{}
	cfg.AddServerConfigLoader(config)
}
