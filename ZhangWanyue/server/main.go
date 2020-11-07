package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
	"flag"
	"os"
	"io"
)
const(
	LineFeed = "\r\n"
)
type Configs map[string]json.RawMessage
var configPath string

type MainConfig struct {
	Port string `json:"port"`
	IP string `json:"IP"`
	ID string `json:"ID"`
	roomID string `json:"roomID"`
	name string `json:"name"`
}

var conf *MainConfig
var confs Configs

var instanceOnce sync.Once

//从配置文件中载入json字符串
func LoadConfig(path string) (Configs, *MainConfig) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicln("load config conf failed: ", err)
	}
	mainConfig := &MainConfig{}
	err = json.Unmarshal(buf, mainConfig)
	if err != nil {
		log.Panicln("decode config file failed:", string(buf), err)
	}
	allConfigs := make(Configs, 0)
	err = json.Unmarshal(buf, &allConfigs)
	if err != nil {
		log.Panicln("decode config file failed:", string(buf), err)
	}

	return allConfigs, mainConfig
}

//初始化 可以运行多次
func SetConfig(path string) {
	allConfigs, mainConfig := LoadConfig(path)
	configPath = path
	conf = mainConfig
	confs = allConfigs
}

// 初始化，只能运行一次
func Init(path string) *MainConfig {
	if conf != nil && path != configPath {
		log.Printf("the config is already initialized, oldPath=%s, path=%s", configPath, path)
	}
	instanceOnce.Do(func() {
		allConfigs, mainConfig := LoadConfig(path)
		configPath = path
		conf = mainConfig
		confs = allConfigs
	})

	return conf
}

//初始化配置文件 为 struct 格式
func Instance() *MainConfig {
	if conf == nil {
		Init(configPath)
	}
	return conf
}

//获取配置文件路径
func ConfigPath() string {
	return configPath
}

func CreateFile(name string) error {
	fo, err := os.Create(name)
	if err != nil {
		return err
	}
	defer func() {
		fo.Close()
	}()
	return nil
}
func WriteLog(filePath, msg string) error {
	if !IsExist(filePath) {
		return CreateDir(filePath)
	}
	var (
		err error
		f   *os.File
	)

	f, err = os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	_, err = io.WriteString(f, LineFeed+msg)

	defer f.Close()
	return err
}
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}
func CreateDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	os.Chmod(path, os.ModePerm)
	return nil
}
var filename = flag.String("config", "./config.json", "Input file path.")

func main() {
	flag.Parse()
	path:= *filename
	fmt.Println("path: ",path)
	Init(path)
	IP := confs["IP"]
	port := confs["port"]
	ID := confs["ID"]
	roomID := confs["roomID"]
	name := confs["name"]
	fmt.Println(string(IP))
	fmt.Println(string(port))
	fmt.Println(string(ID))
	fmt.Println(string(roomID))
	fmt.Println(string(name))

	CreateFile("chat.log")
	WriteLog("chat.log",string(IP) + string(port) + string(ID) + string(roomID) + string(name))
}