package main

import (
	"encoding/json"
	"flag"
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

type UserFromJson struct {
	ID        string    `json:"id"`
	Port	  string    `json:"port"`
	RoomID    string    `json:"room_id"`
	Name      string    `json:"name"`
	IP		  string	`json:"ip"`
}

func main() {
	//利用flag从命令行读取参数
	var config string
	flag.StringVar(&config, "config", "config.json", "读取config.json文件")

	//写入log日志
	logFile, err := os.OpenFile("./chat.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		log.Println("Log file opening error: %v\n", err)
		return
	}
	log.SetOutput(logFile)

	//这里有一个非常重要的操作,转换,必须调用该方法
	flag.Parse()

	//读取json文件
	jsonFile, err := os.Open(config)
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully opened config.json")
	defer jsonFile.Close()
	//同故宫ioutil读取
	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println(err)
		return
	}
	var curUser UserFromJson
	json.Unmarshal(data, &curUser)
	log.Printf("IP: %v\n", curUser.IP)
	log.Printf("Port: %v\n", curUser.Port)
	log.Printf("ID: %v\n", curUser.ID)
	log.Printf("Room_ID: %v\n", curUser.RoomID)
	log.Printf("Name: %v\n", curUser.Name)

	//socket
	c, err := gosocketio.Dial(
		gosocketio.GetUrl("124.236.22.20", 3001, false),
		transport.GetDefaultWebsocketTransport())
	if err != nil {
		log.Printf("websocket connection error: %v\n", err)
		return
	}

	log.Println("connect success\n\n")

	_ = c.On("success", func(c *gosocketio.Channel, v interface{}) {
		log.Println("success:", v)

	})
	//线程
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

	time.Sleep(5 * time.Second)
	log.Println("finished")
}
