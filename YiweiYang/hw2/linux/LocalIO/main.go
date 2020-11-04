package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	config string
)

type User struct {
	IP     string `json:"ip"`
	Port   string `json:"port"`
	ID     string `json:"id"`
	RoomID string `json:"room_id"`
	Name   string `json:"name"`
}

func init() {
	//解析命令行输入
	flag.StringVar(&config, "config", "config.json", "json文件路径，默认为同目录的config.json")
	flag.Parse()
}

func main() {
	readJson(config)
}

func readJson(config string) {
	filePtr, err := os.Open(config)
	if err != nil {
		fmt.Printf("Open file failed [Err:%s]\n", err.Error())
		return
	}
	//延迟
	defer filePtr.Close()

	var userConfig User
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&userConfig)
	if err != nil {
		fmt.Println("Decoder failed", err.Error())
	} else {
		fmt.Printf("IP: %s\n", userConfig.IP)
		fmt.Printf("Port: %s\n", userConfig.Port)
		fmt.Printf("ID: %s\n", userConfig.ID)
		fmt.Printf("Room ID: %s\n", userConfig.RoomID)
		fmt.Printf("Name: %s\n", userConfig.Name)
		writeLog(userConfig)
		fmt.Println("Completed!")
	}
}

func writeLog(userConfig User) {
	//打开写入文件
	logFile, err := os.OpenFile("./chat.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModeAppend|os.ModePerm)
	// 关闭文件
	defer logFile.Close()
	if err != nil {
		panic(err.Error())
	} else {
		logger := log.New(logFile, "[log]", log.LstdFlags|log.Lshortfile|log.LUTC)
		logger.Printf("IP: %s\n", userConfig.IP)
		logger.Printf("Port: %s\n", userConfig.Port)
		logger.Printf("ID: %s\n", userConfig.ID)
		logger.Printf("Room ID: %s\n", userConfig.RoomID)
		logger.Printf("Name: %s\n", userConfig.Name)
	}
}
