package main

import (
	// "flag"
	"fmt"
	// "runtime"
	"github.com/spf13/viper"
	// "bytes"
	// "os"
)

// Config 配置文件的结构
type Config struct {
	Host      host
	Datastore datastore
}

type host struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
	Host    string `json:"host"`
}

type datastore struct {
	Metric    host
	Warehouse host
}

var (
	// Conf 配置文件所在的变量
	Conf     *Config
	confPath string
)

// InitConfig 初始化配置
func InitConfig() (err error) {
	Conf = NewConfig()

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("json")
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file : %s ", err))
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unable to decode into structL %s ", err))
	}
	return nil
}

// NewConfig 生成一个新的配置文件
func NewConfig() *Config {
	return &Config{}
}

func main() {
	InitConfig()
	fmt.Println("hello")
	fmt.Println("host address:", viper.Get("datastore.metric.host"))
	fmt.Println(Conf)
}
