package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

//Actions Constants
const (
	ActionChangeMedia = "change_media" //TODO
	ActionLogin       = "login"
	ActionProcess     = "get_process"
	ActionScreenshot  = "get_screenshot"
	ActionStats       = "get_stats"
)

var addr = flag.String("addr", "0.0.0.0:3567", "http service address") //TODO: Make configuration file and receive two separate parameters (host and port)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	}, //TODO: Delete this in production
}

//Process the received message and perform the actions according to the type
func handleMessage(m []byte, c *websocket.Conn, mt int) {
	var message Message

	if err := json.Unmarshal(m, &message); err != nil {
		log.Println("read:", err)
		return
	}

	responseMessage := &ResponseMessage{
		Action:     message.Action,
		StatusCode: 200,
		Token:      "asdasjkdjksd", //TODO: JWT Token
	}

	var res []byte

	switch message.Action {
	case ActionLogin:
		res, _ = json.Marshal(responseMessage)
	case ActionProcess:
		res = getProcess(message, responseMessage)
	case ActionScreenshot:
		res = getScreenShot(message, responseMessage)
	case ActionStats:
		res = getStats(message, responseMessage)
	}

	if res != nil {
		c.WriteMessage(mt, res)
	}
}

//Upgrade request to websocket connection.
func server(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, rawMessage, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", rawMessage)
		go handleMessage(rawMessage, c, mt)
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", server)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
