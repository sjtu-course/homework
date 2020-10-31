package main

import (
	"os"
	"log"
	"time"
	"strconv"
	"io/ioutil"
	"encoding/json"

	"flag"

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

type Usr struct {
	ID     string `json:"id"`
	Port   string `json:"port"`
	RoomID string `json:"room_id"`
	Name   string `json:"name"`
	IP     string `json:"ip"`
}

func main() {
	// set --config
	var config string
	flag.StringVar(&config, "config", "config.json", "json file")
	logFile, err := os.OpenFile("./chat.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil{
		println("open or create 'chat.log' fail")
		os.Exit(-1)
	}
	log.SetOutput(logFile)

	flag.Parse()

	// read in json
	jsonFile, err := os.Open(config)
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully opened config.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var usr Usr
	json.Unmarshal(byteValue, &usr)

	// log.Println(usr.Port)

	// socketio part
	PORT, err := strconv.Atoi(usr.Port)
	if err != nil {
		log.Println(err)
	}
	c, err := gosocketio.Dial(
		gosocketio.GetUrl(usr.IP, PORT, false),
		transport.GetDefaultWebsocketTransport())
	if err != nil {
		log.Printf("websocket connection error: %v\n", err)
		return
	}

	_ = c.On("success", func(c *gosocketio.Channel, v interface{}) {
		log.Println("success:", v)
	})

	go func(usr Usr) {
		user := User {
						ID:        usr.ID,
						RoomID:    usr.RoomID,
						Name:      usr.Name,
						AvatarURL: "http://q.qlogo.cn/headimg_dl?dst_uin=5684277&spec=100",
						Type:      "user",
		}

		err = c.Emit("login", user)

		if err != nil {
			log.Println("emit error:", err)
		}
	} (usr)

	time.Sleep(5 * time.Second)
	log.Println("finished")
}