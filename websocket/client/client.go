/*
 *  Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
 *
 *  This Source Code Form is subject to the terms of the MIT License.
 *  If a copy of the MIT was not distributed with this file,
 *  You can obtain one at https://github.com/houseme/url-shortenter.
 */

package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	// WebSocket server URL.
	url := "ws://localhost:8080/ws"
	// Create a WebSocket connection.
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Failed to connect to WebSocket server:", err)
	}
	defer func() {
		// Close the WebSocket connection when the main function exits.
		_ = conn.Close()
		log.Println("WebSocket connection closed.")
	}()
	log.Println("WebSocket connection established.")
	// Start a goroutine to handle incoming messages.
	go handleMessagesClient(conn)
	// Send a message to the WebSocket server.
	message := []byte("Hello, WebSocket server!")
	err = conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Println("Failed to send message to WebSocket server:", err)
	}
	// Wait for the user to press Enter.
	log.Println("Press Enter to exit... Exiting.")
	_, _ = fmt.Scanln()
	log.Println("Exiting.")
}
func handleMessagesClient(conn *websocket.Conn) {
	defer func() {
		// Close the WebSocket connection when the handleMessages function exits.
		defer func() {
			_ = conn.Close()
			log.Println("WebSocket connection closed.")
		}()
		log.Println("WebSocket connection closed.")
	}()
	for {
		// Read a message from the WebSocket connection.
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("WebSocket error:", err)
			}
			break
		}
		err = conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("Failed to send message to WebSocket server:", err)
		}
		log.Printf("Message received from WebSocket server: %s", string(message))
	}
}
