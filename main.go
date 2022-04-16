package main

import (
	"encoding/json"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/go-vgo/robotgo"
	"github.com/gorilla/websocket"
)

type Message struct {
	Type string `json:"type"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
}

var upgrader = websocket.Upgrader{}
var lastx = 0
var lasty = 0
var addr = flag.String("addr", "0.0.0.0:8000", "The http service address")
var sens = flag.Int("sensitivity", 100, "The sensitivity of the mouse")
var SENSITIVITY = 0.0
var content, _ = os.ReadFile("templates/index.html")
var homeTemplate = template.Must(template.New("").Parse(string(content)))

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/ws")
}

func ws(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print(err)
		return
	}
	defer c.Close()
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		var msg Message
		json.Unmarshal([]byte(message), &msg)
		if msg.Type == "touchstart" {
			lastx = int(float64(msg.X) / SENSITIVITY)
			lasty = int(float64(msg.Y) / SENSITIVITY)
		} else if msg.Type == "touchmove" {
			robotgo.MoveRelative(int(float64(msg.X)/SENSITIVITY)-lastx, int(float64(msg.Y)/SENSITIVITY)-lasty)
			lastx = int(float64(msg.X) / SENSITIVITY)
			lasty = int(float64(msg.Y) / SENSITIVITY)
		} else if msg.Type == "touchend" {
			lastx = 0
			lasty = 0
		}
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	SENSITIVITY = 1.0 / (float64(*sens*2) / 100.0)
	log.Printf("Running http server on http://%s", *addr)
	log.Println("Sensitivity:", *sens)
	http.HandleFunc("/", home)
	http.HandleFunc("/ws", ws)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
