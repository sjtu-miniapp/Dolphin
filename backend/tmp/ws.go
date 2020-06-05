package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func hi(w http.ResponseWriter, r *http.Request) {
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	c, err := upGrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("upgrade: %s", err)
		return
	}

	defer c.Close()
	errChan := make(chan error)
	go func() {
		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				errChan <- err
				break
			}
			if message == nil {
				errChan <- nil
				break
			}
		}
		// write to mq: openid/task_id/version/string
	}()
	// use a conditional value
	go func() {
		for {
			// fetch from mq to know about
			// fetch from mongo
			err = c.WriteMessage(mt, message)
			if err != nil {
				log.Println("write:", err)
				errChan <- err
				break
			}
		}

	}()
	<- errChan
}

//func main() {
//
//}
