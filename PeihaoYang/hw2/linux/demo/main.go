package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

// User data
type User struct {
	ID        string    `json:"id"`
	RoomID    string    `json:"room_id"`
	Name      string    `json:"name"`
	AvatarURL string    `json:"avatarUrl"`
	Type      string    `json:"type"`
	Address   string    `json:"address"`
	LoginTime time.Time `json:"login_time"`
}

// 从config.json配置读取IP，port，ID，roomID，name信息
// 命令行传入config.json参数，即 ./main -config config.json
// 将日志写入文件系统，chat.log

func loadJSON(filepath string) ([]byte, error) {
	content := make([]byte, 300)
	configFile, err := os.OpenFile(filepath, os.O_RDONLY, os.ModePerm)
	defer configFile.Close()
	if err != nil {
		fmt.Println("Cannot open config file.")
		return content, err
	}
	n, err := configFile.Read(content)
	if err != nil {
		fmt.Println("Cannot read the content.")
		return content, err
	}
	return content[:n], nil
}

func main() {
	argConfig := flag.String("config", "config", "Config file name")
	argLog := flag.String("log", "chat", "Log file name")
	flag.Parse()
	configName := fmt.Sprintf("%s.json", *argConfig)
	logName := fmt.Sprintf("%s.log", *argLog)
	fmt.Println("config input file:", configName)
	fmt.Println("log output file:", logName)

	logFile, err := os.Create(logName)
	if err != nil {
		fmt.Println("Cannot create log file.")
		return
	}
	defer logFile.Close()

	chatLog := log.New(logFile, "[Log]", log.Lmicroseconds|log.Llongfile)

	configJSON, err := loadJSON(configName)
	if err != nil {
		chatLog.Printf("json loading error: %v\n", err)
		return
	}
	userPtr := &User{}
	json.Unmarshal(configJSON, userPtr)

	c, err := gosocketio.Dial(
		gosocketio.GetUrl("124.236.22.20", 3001, false),
		transport.GetDefaultWebsocketTransport())
	if err != nil {
		chatLog.Printf("websocket connection error: %v\n", err)
		return
	}

	chatLog.Printf("connect success\n")

	_ = c.On("success", func(c *gosocketio.Channel, v interface{}) {
		chatLog.Println("success:", v)
	})

	go func() {
		err = c.Emit("login", userPtr)
		if err != nil {
			chatLog.Println("emit error:", err)
		}
	}()

	time.Sleep(5 * time.Second)
	chatLog.Println("finished")
}
