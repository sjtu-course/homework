package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
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

//定义配置文件解析后的结构
type Config struct {
	IP     string `json:"ip"`
	Port   string `json:"port"`
	ID     string `json:"ID"`
	RoomID string `json:"roomID"`
	Name   string `json:"name"`
}

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (jst *JsonStruct) Load(filename string, v interface{}) {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}


func init() {
	file := "./" + "chat" + ".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	loger = log.New(logFile, "[qSkiptool]",log.LstdFlags | log.Lshortfile | log.LUTC) // 将文件设置为loger作为输出
	return
}
var loger *log.Logger
func main() {

	var src=flag.String("config", "", "config file")
	flag.Parse()
	fmt.Println(*src)
	JsonParse := NewJsonStruct()
	v := Config{}
	JsonParse.Load(*src, &v)
	fmt.Println(v.Port)
	fmt.Println(v.IP)
	port, err := strconv.Atoi(v.Port)
	c, err := gosocketio.Dial(
		gosocketio.GetUrl(v.IP, port, false),
		transport.GetDefaultWebsocketTransport())
	if err != nil {
		loger.Printf("websocket connection error: %v\n", err)
		return
	}

	loger.Printf("connect success\n\n")

	_ = c.On("success", func(c *gosocketio.Channel, v interface{}) {
		loger.Println("success:", v)
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
			loger.Println("emit error:", err)
		}
	}()

	time.Sleep(5*time.Second)
	loger.Println("finished")
}
