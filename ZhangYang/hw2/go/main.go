package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"fmt"
)
var result map[string]interface{}
func init() {
	file := "./" +"chat"+ ".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
	log.SetPrefix("Chat")
	log.SetFlags(log.LstdFlags | log.Lshortfile |log.Ldate |log.Ltime)
	return
}
func main() {
	var (
		config string
		h bool
	)
	flag.StringVar(&config,"config","","read json info from the json file you give")
	flag.BoolVar(&h,"h",false,"usage of this command")
	flag.Parse()
	if h{
		flag.PrintDefaults()
	}
	if config!=""{
		readJson(config)
	}
}
func readJson(configPath string)  {
	// 打开json文件
	jsonFile, err := os.Open(configPath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	fmt.Println(string(byteValue))
	log.Println("readJson")
	log.Println(string(byteValue))
	json.Unmarshal([]byte(byteValue), &result)
}
