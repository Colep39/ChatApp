package main

import (
    "fmt"
	"log"
    "net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,

	// temporarily allowing any connection
	CheckOrigin: func(r *http.Request) bool {return true},
}

// reader that listens for new messages being sent to the websocket end point
func reader(conn *websocket.Conn){
	for {
		// reading in the message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// printing out message
		fmt.Println(string(p))
		
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

// defining our websocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.Host)
	// upgrading the connection to a web socket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil{
		log.Println(err)
	}
	reader(ws)
}

func setupRoutes() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Simple Server")
    })
	// map the /ws endpoint to the serveWs function
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Chat App v0.01")
    setupRoutes()
    http.ListenAndServe(":8080", nil)
}