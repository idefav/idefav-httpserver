package env

import (
	"idefav-httpserver/cfg"
	"math"
	"os"
	"strconv"
)

const (
	APPNAME                        = "APP-ID"
	ADDRESS                        = "ADDR"
	SHUTDOWN_GRACEFUL              = "GRACEFUL"
	SHUTDOWN_GRACEFUL_WAIT_TIME_MS = "GRACEFUL-WAIT-TIME-MS"
	WARMUP                         = "WARMUP"
	ROUTER_NAME                    = "ROUTER-NAME"
	ACCESS_LOG                     = "ACCESS-LOG"
	CONFIG_FILE                    = "CONFIG"
)

type Environment struct {
	Name  string
	Order int
}

func (e Environment) GetName() string {
	return e.Name
}

func (e Environment) GetOrder() int {
	return e.Order
}

func (e Environment) Load(config *cfg.ServerConfig) {

	if val := GetStringEnvVal(APPNAME); val != "" {
		config.AppName = val
	}

	if val := GetStringEnvVal(ADDRESS); val != "" {
		config.Address = val
	}

	if val, err := GetBoolEnvVal(SHUTDOWN_GRACEFUL); err == nil {
		config.GracefulShutdown = val
	}

	if val, err := GetIntEnvVal(SHUTDOWN_GRACEFUL_WAIT_TIME_MS); err == nil {
		config.GracefulShutdownWaitTimeMs = val
	}

	if val, err := GetBoolEnvVal(WARMUP); err == nil {
		config.Warmup = val
	}

	if val := GetStringEnvVal(ROUTER_NAME); val != "" {
		config.RouterName = val
	}

	if val, err := GetBoolEnvVal(ACCESS_LOG); err == nil {
		config.AccessLog = val
	}

	if val := GetStringEnvVal(CONFIG_FILE); val != "" {
		config.ConfigFile = val
	}

}

func GetStringEnvVal(key string) string {
	val := os.Getenv(key)
	return val
}

func GetBoolEnvVal(key string) (bool, error) {
	val := GetStringEnvVal(key)
	result, err := strconv.ParseBool(val)
	return result, err
}

func GetIntEnvVal(key string) (int, error) {
	val := GetStringEnvVal(key)
	result, err := strconv.ParseInt(val, 10, 0)
	return int(result), err
}

func init() {
	environment := &Environment{
		Name:  "Environment",
		Order: math.MaxInt - 1,
	}
	cfg.AddServerConfigLoader(environment)
}
