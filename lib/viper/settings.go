package main

import (
	"fmt"
	"github.com/spf13/viper"
)

var dbConf = new(MySqlConf)

type MySqlConf struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func way1() {
	viper.SetConfigFile("./config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("viper.ReadInConfig failed, err:%v\n", err)
		return
	}

	if err := viper.Unmarshal(dbConf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		return
	}

	fmt.Printf("dbConf:%#v\n", dbConf)

}

func way2() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("viper.ReadInConfig failed, err:%v\n", err)
		return
	}

	if err := viper.Unmarshal(dbConf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		return
	}

	fmt.Printf("dbConf:%#v\n", dbConf)
}

func main() {
	//way1()
	way2()
}
