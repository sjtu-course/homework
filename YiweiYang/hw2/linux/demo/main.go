package main

import (
	"fmt"
	"strings"
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

func main() {
	c, err := gosocketio.Dial(
		gosocketio.GetUrl("124.236.22.20", 3001, false),
		transport.GetDefaultWebsocketTransport())
	if err != nil {
		fmt.Printf("websocket connection error: %v\n", err)
		return
	}

	fmt.Printf("connect success\n\n")

	_ = c.On("success", func(c *gosocketio.Channel, v interface{}) {
		fmt.Println("success:", v)
		// Analyzing
		m := v.([]interface{})[0]
		data := m.(map[string]interface{})
		adress := data["address"].(string)
		adressArray := strings.Split(adress, ":")
		fmt.Printf("IP: %s\n", adressArray[0])
		fmt.Printf("Port: %s\n", adressArray[1])
		fmt.Printf("ID: %s\n", data["id"].(string))
		fmt.Printf("Room ID: %s\n", data["room_id"].(string))
		fmt.Printf("Name: %s\n", data["name"].(string))
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
			fmt.Println("emit error:", err)
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("finished")
}
