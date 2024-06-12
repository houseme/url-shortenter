/*
 *  Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
 *
 *  This Source Code Form is subject to the terms of the MIT License.
 *  If a copy of the MIT was not distributed with this file,
 *  You can obtain one at https://github.com/houseme/url-shortenter.
 */

package main

import (
    "log"
    "net/http"
    "strconv"
    
    "github.com/gogf/gf/v2/util/grand"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        // Allow all connections.
        return true
    },
}

type Connection struct {
    ID     string
    Socket *websocket.Conn
}

var connections = make(map[string]*Connection)

func main() {
    http.HandleFunc("/ws", handleWebsocketServer)
    
    log.Println("Server started.")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleWebsocketServer(w http.ResponseWriter, r *http.Request) {
    // Upgrade the HTTP connection to a WebSocket connection.
    socket, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Failed to upgrade connection:", err)
        return
    }
    
    // Generate a unique ID for the connection.
    id := grand.Letters(32)
    
    // Create a new connection object.
    conn := &Connection{
        ID:     id,
        Socket: socket,
    }
    
    // Add the connection to the global connections map.
    connections[id] = conn
    
    // Start a goroutine to handle incoming messages.
    go handleMessages(conn)
    message := []byte("Hello, WebSocket Client!")
    for i := 0; i < 10; i++ {
        err = conn.Socket.WriteMessage(websocket.TextMessage, message)
        if err != nil {
            log.Println("Failed to send message to WebSocket server:", err)
        }
        log.Println("Message sent to WebSocket server:", strconv.FormatInt(int64(i), 10))
    }
    log.Println("WebSocket connection established with ID:", id)
}

func handleMessages(conn *Connection) {
    defer func() {
        // Close the WebSocket connection when the handleMessages function exits.
        _ = conn.Socket.Close()
        log.Println("WebSocket connection closed with ID:", conn.ID)
    }()
    
    for {
        // Read a message from the WebSocket connection.
        _, message, err := conn.Socket.ReadMessage()
        if err != nil {
            if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
                log.Println("WebSocket error:", err)
            }
            break
        }
        // message = []byte("Hello, WebSocket Client!")
        err = conn.Socket.WriteMessage(websocket.TextMessage, message)
        if err != nil {
            log.Println("Failed to send message to WebSocket server:", err)
        }
        
        log.Printf("Message received from WebSocket connection %s: %s", conn.ID, string(message))
    }
}
