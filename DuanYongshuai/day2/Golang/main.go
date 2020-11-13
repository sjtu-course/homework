package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

type User struct {
	ID        string    `json:"id"`
	RoomID    string    `json:"room_id"`
	Name      string    `json:"name"`
	AvatarURL string    `json:"avatarUrl"`
	Type      string    `json:"type"`
	Address   string    `json:"address"`
	LoginTime time.Time `json:"login_time"`
}

type Config struct {
	ID			string	`json:"id"`
	IP			string	`json:"ip"`
	Port		int		`json:"port"`
	RoomID		string	`json:"room_id"`
	Name		string	`json:"name"`
}


func main() {
	// set log file
	file, err := os.OpenFile("chat.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("open log file error: %v\n", err)
	}
	log.SetOutput(file)

	// read config.json
	filePath := flag.String("config", "config.json", "filename of config information")
	flag.Parse()

	fileData, err := ioutil.ReadFile(*filePath)
	if err != nil {
		log.Fatalf("open json file error: %v\n", err)
	}

	conf := &Config{}
	json.Unmarshal([]byte(fileData), &conf)
	fmt.Println("*** unmarshal result: ***")
	fmt.Println(conf)

	log.Println("config json load sucess")

	// connect to server
	c, err := gosocketio.Dial(
		gosocketio.GetUrl("124.236.22.20", 3001, false),
		transport.GetDefaultWebsocketTransport())
	if err != nil {
		log.Fatalf("websocket connection error: %v\n", err)
	}

	log.Printf("connect success\n\n")

	_ = c.On("success", func(c *gosocketio.Channel, v interface{}) {
		log.Println("success:", v)
	})

	go func() {
		err = c.Emit("login", User{
			ID:        "user",
			RoomID:    "roomid",
			Name:      "rtcname0",
			AvatarURL: "http://q.qlogo.cn/headimg_dl?dst_uin=5684277&spec=100",
			Type:      "user",
		})

		if err != nil {
			log.Println("emit error:", err)
		}
	}()

	time.Sleep(5*time.Second)
	log.Println("finished")
}
