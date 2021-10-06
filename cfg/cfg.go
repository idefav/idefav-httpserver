package cfg

import "flag"

const (
	OK            = "OK"
	HEALTH        = "Health"
	UNHEALTHY     = "Unhealthy"
	VERSION       = "VERSION"
	ERROR_HANDLER = "ERROR"
)

type ServerConfig struct {
	// listen Address
	Address string
}

// cfg setup
func SetUp() *ServerConfig {
	config := ServerConfig{
		Address: ":8080",
	}
	flag.StringVar(&config.Address, "Address", ":8080", "设置服务监听地址")
	flag.Parse()
	return &config
}
