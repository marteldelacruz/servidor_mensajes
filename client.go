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
}

//
func sendMessage(id string) {
	conn, err := net.Dial(_PROTOCOL, _PORT)

	if err != nil {
		fmt.Println(err)
		return
	}

	// send client ID
	conn.Write([]byte(id + "|message|Oh no"))
}

// The main menu contains all the available options
// for the client
func mainMenu(conn net.Conn) {
	var opt string

	for {
		fmt.Println("1-. Send message")
		fmt.Println("2.- Send file")
		fmt.Println("3.- View chat")
		fmt.Println("0.- Exit")

		fmt.Print("Select an option: ")
		fmt.Scanln(&opt)

		switch opt {
		case "1":
			break
		case "2":
			break
		case "3":
			break
		case "0":
			return
			break
		default:
		}
	}
}

func main() {
	var clientID string

	fmt.Print("Nickname: ")
	fmt.Scanln(&clientID)

	go client(clientID)

	fmt.Scanln()

	go sendMessage(clientID)

	fmt.Scanln()
}
