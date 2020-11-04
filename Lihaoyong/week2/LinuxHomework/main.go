package main

import (
	"fmt"
	"time"
	"os"
	"encoding/json"
	"io/ioutil"
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

func dropErr(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {
	
	//all output go to chat.log
	f, _ := os.OpenFile("chat.log", os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND,0755)
    os.Stdout = f
    os.Stderr = f

	//situation 1
	
		var config = flag.String("config","","-config: Input file")

		flag.Parse()
		fmt.Println("-configFile:", *config)

		//read in filepath
		filePath := *config
		f, err := os.Open(filePath)
		dropErr(err)

		defer f.Close()
		//read from json
		byteValue, _ := ioutil.ReadAll(f)

		var user User
		//read to be struct
		json.Unmarshal([]byte(byteValue), &user)
		//fmt.Println(user.ID)
		


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
		})

		
		go func(_id, _roomID, _name, _avatarURL, _type string) {
			/*
			err = c.Emit("login", User{
				ID:        "user",
				RoomID:    "roomid",
				Name:      "rtcname0",
				AvatarURL: "http://q.qlogo.cn/headimg_dl?dst_uin=5684277&spec=100",
				Type:      "user",
			})
			*/
			
			err = c.Emit("login", User{
				ID:        _id,
				RoomID:    _roomID,
				Name:      _name,
				AvatarURL: _avatarURL,
				Type:      _type,
			})


			if err != nil {
				fmt.Println("emit error:", err)
			}
		}(user.ID, user.RoomID, user.Name, user.AvatarURL, user.Type)

		time.Sleep(5*time.Second)
		fmt.Println("finished")

	
}
