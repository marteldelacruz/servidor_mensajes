package main

import (
	"fmt"
	"net"
)

var _PROTOCOL = "tcp"
var _PORT = ":9999"

// The new client will connect to a server on the PORT
// if the server is not running already, an error message
// will be shown
func client(id string) {
	conn, err := net.Dial(_PROTOCOL, _PORT)

	if err != nil {
		fmt.Println(err)
		return
	}
	// send client ID
	conn.Write([]byte(id))

	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	var clientID string

	fmt.Print("Nickname: ")
	fmt.Scanln(&clientID)

	go client(clientID)
	fmt.Scanln()
}
