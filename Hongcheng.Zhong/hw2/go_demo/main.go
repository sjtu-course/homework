package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

// User for login
type User struct {
	ID        string    `json:"id"`
	RoomID    string    `json:"room_id"`
	Name      string    `json:"name"`
	AvatarURL string    `json:"avatarUrl"`
	Type      string    `json:"type"`
	Address   string    `json:"address"`
	LoginTime time.Time `json:"login_time"`
}

// Config to use
type Config struct {
	IP     string `json:"IP"`
	Port   int    `json:"Port"`
	ID     string `json:"ID"`
	RoomID string `json:"RoomID"`
	Name   string `json:"Name"`
}

// log related
var (
	Info *log.Logger
	// Warning *log.Logger
	Error *log.Logger
)

func init() {
	logFile, err := os.OpenFile("chat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("fail to open logfile：", err)
	}
	Info = log.New(io.MultiWriter(os.Stdout, logFile), "Info:", log.Ldate|log.Ltime|log.Lshortfile)
	// Warning = log.New(io.MultiWriter(os.Stdout, logFile), "Warning:", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr, logFile), "Error:", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	var confPath string
	flag.StringVar(&confPath, "config", "config.json", "path to config json file")
	flag.Parse()

	configFile, err := os.Open(confPath)
	if err != nil {
		Error.Printf("config file error: %s\n", err)
		return
	}

	var conf = &Config{}
	err = json.NewDecoder(configFile).Decode(conf)
	if err != nil {
		Error.Printf("Failed to decode %s, %s\n", "config.json", err)
	}
	Info.Printf("config file: %s is parsed successfully", confPath)

	c, err := gosocketio.Dial(
		gosocketio.GetUrl(conf.IP, conf.Port, false),
		transport.GetDefaultWebsocketTransport())
	if err != nil {
		Error.Printf("websocket connection error: %v\n", err)
		return
	}

	Info.Printf("connect success\n\n")

	_ = c.On("success", func(c *gosocketio.Channel, v interface{}) {
		Info.Println("success:", v)
	})

	go func() {
		err = c.Emit("login", User{
			ID:        conf.ID,
			RoomID:    conf.RoomID,
			Name:      conf.Name,
			AvatarURL: "http://q.qlogo.cn/headimg_dl?dst_uin=5684277&spec=100",
			Type:      "user",
		})

		if err != nil {
			fmt.Println("emit error:", err)
		}
	}()

	time.Sleep(5 * time.Second)
	Info.Println("finished")
}
