package main

import (
	"encoding/json"
	"errors"
	"flag"
	"io"
	"log"
	"os"
	"time"

	gosocketio "github.com/graarh/golang-socketio"

	"github.com/graarh/golang-socketio/transport"
)

func writeConfig(path string, v Config) (int, error) {
	fp, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	data, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}

	n, err := fp.Write(data)
	if err != nil {
		log.Fatal(err)
	}

	return n, err
}

func readConfig(path string) (Config, error) {
	_, err := os.Stat(path)
	if err != nil {
		if !os.IsExist(err) {
			return Config{}, errors.New("No such config file")
		}
	}
	fp, err := os.OpenFile(path, os.O_RDONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	data := make([]byte, 1024)
	n, err := fp.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	config := Config{}
	err = json.Unmarshal(data[:n], &config)
	if err != nil {
		log.Fatal(err)
	}
	return config, err
}

// User
type User struct {
	ID        string    `json:"id"`
	RoomID    string    `json:"room_id"`
	Name      string    `json:"name"`
	AvatarURL string    `json:"avatarUrl"`
	Type      string    `json:"type"`
	Address   string    `json:"address"`
	LoginTime time.Time `json:"login_time"`
}

// Config
type Config struct {
	IP     string `json:"ip"`
	Port   int    `json:"port"`
	ID     string `json:"id"`
	RoomID string `json:"roomId"`
	Name   string `json:"name"`
}

func main() {
	// 配置logger
	logFile, _ := os.OpenFile("./chat.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	defer logFile.Close()
	writers := []io.Writer{logFile, os.Stdout} // 同时输出到日志文件和Stdout
	logger := log.New(io.MultiWriter(writers...), "[RTC]", log.Lshortfile|log.Ltime|log.Ldate)

	// 获取配置文件路径
	configPath := flag.String("config", "./config.json", "The path of the config file")
	flag.Parse()

	// 读取配置文件
	config, err := readConfig(*configPath)
	if err != nil {
		// 若配置文件不存在，使用默认配置
		config = Config{
			IP:     "124.236.22.20",
			Port:   3001,
			RoomID: "default_room_id",
			ID:     "default_id",
			Name:   "default_user_name",
		}
		writeConfig(*configPath, config)
	}

	c, err := gosocketio.Dial(
		gosocketio.GetUrl(config.IP, config.Port, false),
		transport.GetDefaultWebsocketTransport())
	if err != nil {
		logger.Printf("websocket connection error: %v\n", err)
		return
	}

	logger.Printf("connect success\n\n")

	_ = c.On("success", func(c *gosocketio.Channel, v interface{}) {
		logger.Println("success:", v)
	})

	// 根据配置文件调用login接口
	go func(config Config) {
		user := User{
			ID:        config.ID,
			RoomID:    config.RoomID,
			Name:      config.Name,
			AvatarURL: "http://q.qlogo.cn/headimg_dl?dst_uin=5684277&spec=100",
			Type:      "user",
		}
		err = c.Emit("login", user)

		if err != nil {
			logger.Println("emit error:", err)
		}
	}(config)

	time.Sleep(5 * time.Second)
	logger.Println("finished")
}
