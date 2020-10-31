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

type User struct {
	ID        string    `json:"id"`
	RoomID    string    `json:"room_id"`
	Name      string    `json:"name"`
	AvatarURL string    `json:"avatarUrl"`
	Type      string    `json:"type"`
	Address   string    `json:"address"`
	LoginTime time.Time `json:"login_time"`
}

func readFile(jsonpath string) []byte {
	jsonfile, err := os.Open(jsonpath)
	if err != nil {
		fmt.Println("Open file failed [Err:%s]", err.Error())
	}
	defer jsonfile.Close()
	jsonconfig := make([]byte, 512)
	n, err := jsonfile.Read(jsonconfig)
	if err != nil {
		fmt.Println("Read file failed [Err:%s]", err.Error())
	}
	return jsonconfig[:n]
}

func main() {
	jsonpath := flag.String("", "config.json", "jsonname")
	logname := "./chat.log"
	logfile, err := os.Create(logname)
	if err != nil {
		fmt.Println("Create log failed [Err:%s]", err.Error())
		return
	}
	defer logfile.Close()
	chatlog := log.New(logfile, "[ChatLog] ", log.Llongfile) //New(out io.Writer, prefix string, flag int)
	configjson := readFile(*jsonpath)
	flag.Parse()
	user := &User{}
	json.Unmarshal(configjson, user)

	c, err := gosocketio.Dial(
		gosocketio.GetUrl("124.236.22.20", 3001, false),
		transport.GetDefaultWebsocketTransport())
	if err != nil {
		fmt.Printf("websocket connection error: %v\n", err)
		chatlog.Printf("websocket connection error: %v\n", err)
		return
	}

	fmt.Printf("connect success\n\n")
	chatlog.Printf("connect success\n\n")

	_ = c.On("success", func(c *gosocketio.Channel, v interface{}) {
		fmt.Println("success:", v)
		chatlog.Println("success:", v)
	})

	go func() {
		err = c.Emit("login", user)
		if err != nil {
			fmt.Println("emit error:", err)
			chatlog.Println("emit error:", err)
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("finished")
	chatlog.Println("finished")
}
