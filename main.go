package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

var members []*websocket.Conn

func main() {
	http.Handle("/chat", websocket.Handler(chat))
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./public"))))
	if e := http.ListenAndServe(":9090", nil); e != nil {
		fmt.Println(e)
	}
}
func chat(ws *websocket.Conn) {
	defer ws.Close()
	members = append(members, ws)
	for {
		str := ""
		if e := websocket.Message.Receive(ws, &str); e != nil {
			break
		}
		go sendToAll(str)
	}
}
func sendToAll(str string) {
	for k, v := range members {
		if e := websocket.Message.Send(v, str); e != nil {
			members = append(members[:k], members[k+1:]...) //delete offline Conn in member
			break
		}
	}
}
