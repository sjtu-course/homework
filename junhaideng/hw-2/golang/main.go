package main 

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"fmt"
)

// config 命令行参数，可用于指定对应的配置文件路径
var config string 

// 初始化的时候做一些配置工作
func init(){
	flag.StringVar(&config,"config" ,"config.json", "the path to configuration json file")
	f, err := os.OpenFile("chat.log", os.O_APPEND|os.O_CREATE, 0666)
	if err != nil{
		println("open or create 'chat.log' fail")
		os.Exit(-1)
	}
	log.SetOutput(f)
}

// Config 类型用来反序列json文件
type Config struct{
	IP string `json:"IP"`
	Port string `json:"port"`
	ID string `json:"ID"`
	RoomID string `json:"roomID"`
	Name string `json:"name"`
}

func main(){
	// 解析参数
	flag.Parse()
	// 打开配置文件
	file, err := os.Open(config)
	if err != nil{
		log.Panicf("open %s failed: %s\n", config, err.Error())
		return 
	}
	// conf 用来接收数据
	var conf = &Config{}
	err = json.NewDecoder(file).Decode(conf)
	if err != nil{
		log.Panicf("decode configuration from %s failed: %s", config, err.Error())
		return 
	}
	// 终端输出，这里仅仅用于展示
	fmt.Printf("%#v", conf)
}