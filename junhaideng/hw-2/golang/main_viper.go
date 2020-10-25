package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

var config string

func init() {
	flag.StringVar(&config, "config", "config.json", "config filename")
	logFile, err := os.OpenFile("chat.log", os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-1)
	}
	logrus.SetOutput(logFile)
	logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	flag.Parse()
	// 设置文件名
	viper.SetConfigFile(config)

	// 初始化
	err := viper.ReadInConfig()
	if err != nil {
		logrus.WithField("main", err.Error()).Errorln("初始化配置文件失败")
		return
	}

	// 下面的可以直接使用调用获取
	fmt.Println(viper.GetString("IP"))
	fmt.Println(viper.GetString("roomID"))
}
