package main

import (
	"encoding/json"
	"errors"
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

func ReadConfig() (User, error) {
	filePath := "config.json"
	user := User{}
	fp, err := os.OpenFile(filePath, os.O_RDONLY, 0755)
	if err != nil {
		fmt.Println("Open file error", err)
		return user, errors.New("Failed to open file.")
	}
	defer fp.Close()
	data := make([]byte, 500)
	n, err := fp.Read(data)
	if err != nil {
		fmt.Println("Read file error:", err)
		return user, errors.New("Failed to read file.")
	}
	json.Unmarshal(data[:n], user)

	return user, nil
}

func main() {
	// create log file
	logFileName := "chat.log"
	logFile, err := os.Create(logFileName)
	if err != nil {
		fmt.Println("Failed to create logging file \"chat.log\"")
		return
	}
	defer logFile.Close()
	chatLog := log.New(logFile, "[Chat]", log.Llongfile)

	client, err := gosocketio.Dial(
		gosocketio.GetUrl("124.236.22.20", 3001, false),
		transport.GetDefaultWebsocketTransport())
	if err != nil {
		chatLog.Printf("websocket connection error: %v\n", err)
		return
	}

	chatLog.Printf("connect success\n\n")

	_ = client.On("success", func(c *gosocketio.Channel, v interface{}) {
		chatLog.Println("success:", v)
	})

	go func() {
		user, err := ReadConfig()
		if err != nil {
			chatLog.Println("Read config error:", err)
			return
		}
		err = client.Emit("login", user)
		if err != nil {
			chatLog.Println("emit error:", err)
		}
	}()

	time.Sleep(5*time.Second)
	chatLog.Println("finished")
}
