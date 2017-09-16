// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"sync"
	"log"
	"net/url"
	"os"
	"os/signal"
	"fmt"
	"github.com/gorilla/websocket"
)

var wg sync.WaitGroup


func main() {
	
	addr := "echo.websocket.org"
	interrupt := make(chan os.Signal, 1)

	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: addr }

	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	wg.Add(1)
	
	// Listen
	go listenToWS(conn)

	// Send something
	sendToWS(conn, "Woot woo")
	sendToWS(conn, "Another pile of shite")
	sendToWS(conn, "And again...")
	
	wg.Wait()	

	defer conn.Close()
	
}

func sendToWS(conn *websocket.Conn, s string) {
	fmt.Println("\nSending message:",s)
	conn.WriteMessage(websocket.TextMessage, []byte(s))	
}


func listenToWS(conn *websocket.Conn) {

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", message)
	}

}
