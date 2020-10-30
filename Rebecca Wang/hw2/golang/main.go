package main

import (
	"os"
	"time"
	"flag"
	"encoding/json"
	"log"

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

func testRead(filepath string) []byte {
	fp, err := os.OpenFile(filepath, os.O_RDONLY, 0755)
	defer fp.Close()
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 500)
	n, err := fp.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	return data[:n]
}

func main() {
	fileName := "chat.log"
	logFile, errlog := os.Create(fileName)
	defer logFile.Close()
	if errlog != nil {
		log.Fatalln("open file error !")
	}
	chatLog := log.New(logFile, "[Chat]", log.Llongfile)

	config := flag.String("config", "", "Config info")
	flag.Parse()

	var data []byte
	data = testRead(*config)
	user := &User{}
	json.Unmarshal([]byte(data), user)

	c, err := gosocketio.Dial(
		gosocketio.GetUrl("124.236.22.20", 3001, false),
		transport.GetDefaultWebsocketTransport())
	if err != nil {
		chatLog.Printf("websocket connection error: %v\n", err)
		return
	}

	chatLog.Printf("connect success\n\n")

	_ = c.On("success", func(c *gosocketio.Channel, v interface{}) {
		chatLog.Println("success:", v)
	})

	go func() {
		err = c.Emit("login", user)

		if err != nil {
			chatLog.Println("emit error:", err)
		}
	}()

	time.Sleep(5*time.Second)
	chatLog.Println("finished")
}
