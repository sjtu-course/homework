package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	configFile, err := os.Open("config.json")
	if err != nil {
		fmt.Printf("config file error: %s\n", err)
		return
	}

	var conf = &Config{}
	err = json.NewDecoder(configFile).Decode(conf)
	if err != nil {
		fmt.Printf("Failed to decode %s, %s\n", "config.json", err)
	}

	c, err := gosocketio.Dial(
		gosocketio.GetUrl(conf.IP, conf.Port, false),
		transport.GetDefaultWebsocketTransport())
	if err != nil {
		fmt.Printf("websocket connection error: %v\n", err)
		return
	}

	fmt.Printf("connect success\n\n")

	_ = c.On("success", func(c *gosocketio.Channel, v interface{}) {
		fmt.Println("success:", v)
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
	fmt.Println("finished")
}
