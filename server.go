package main

import (
	"fmt"
	"net"
)

var PROTOCOL = "tcp"
var PORT = ":9999"

// Verifies if a client adress already exist on the list
func clientIsInList(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

/// This rutine runs the server on a loop to keep
/// handling client petitions using the TCP connection
/// on the 9999 port
func server() {
	server, err := net.Listen(PROTOCOL, PORT)
	var adressList []string

	// terminate when an error ocurrs
	if err != nil {
		fmt.Println(err)
		return
	}

	// loop to handle clients
	for {
		client, err := server.Accept()

		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleClient(client, &adressList)
	}
}

// This function takes charge of handling clients
// by sending them a process the first time they
// connect to the server
func handleClient(client net.Conn, clientsList *[]string) {
	// wait for client ID
	data := make([]byte, 100)
	br, err := client.Read(data)

	// terminate when an error ocurrs
	if err != nil {
		fmt.Println(err)
		return
	} else {
		clientID := string(data[:br])
		// check if client already exist
		if clientIsInList(clientID, *clientsList) {
		} else {
			fmt.Println(clientID + " connected...")
			*clientsList = append(*clientsList, clientID)
		}
	}

}

func main() {
	go server()

	// press enter to exit...
	fmt.Scanln()
}
