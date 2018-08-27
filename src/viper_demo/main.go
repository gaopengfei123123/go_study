package main

import (
	// "flag"
	"fmt"
	// "runtime"
	"github.com/spf13/viper"
	// "bytes"
	// "os"
)


type Config struct{
	Host host
	Datastore datastore
}

type host struct{
	Address string `json:"address"`
	Port	int `json:"port"`
	Host	string `json:"host"`
}

type datastore struct{
	Metric 		host	
	Warehouse 	host
}

var (
	Conf  *Config
	confPath string
)

func InitConfig() (err error){
	Conf = NewConfig()

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("json")
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file : %s \n", err))
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unable to decode into structL %s \n", err))
	}
	return nil
}

func NewConfig() *Config{
	return &Config{}
}

func main(){
	InitConfig()
	fmt.Println("hello")
	fmt.Println("host address:", viper.Get("datastore.metric.host"))
	fmt.Println(Conf)
}