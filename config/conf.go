package config

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

var Config *AppConfig

type ServerConfig struct {
	DeviceName string `mapstructure:"device_name"`
	SnapLen    int32  `mapstructure:"snap_len"`
	Promisc    bool
	Timeout    time.Duration
	BpfFilter  string `mapstructure:"bpf_filter"`
}

type RoutineConfig struct {
	MonitorRoutine bool          `mapstructure:"monitor_routine"`
	CheckInterval  time.Duration `mapstructure:"check_interval"`
}

type LogConfig struct {
	ErrorLog string `mapstructure:"error_log"`
	InfoLog  string `mapstructure:"info_log"`
}

type AppConfig struct {
	Server  ServerConfig  `mapstructure:"server"`
	Log     LogConfig     `mapstructure:"log"`
	Routine RoutineConfig `mapstructure:"routine"`
}

type WriteConfig struct {
	Writing bool
}

func init() {
	viper.SetConfigType("toml")
	viper.SetConfigFile("./config/config.toml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}
	Config = NewAppConfig()
	if err := viper.Unmarshal(Config); err != nil {
		log.Fatalf("Error unmarshaling config: %v", err)
	}
	fmt.Printf("解析配置文件,%+v\n", Config)
}

func NewAppConfig() *AppConfig {
	Config = &AppConfig{}
	Config.Server.SnapLen = 1024
	Config.Server.Promisc = false
	return Config
}
