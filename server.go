package main

import (
	"fmt"
	"net"
	"strings"

	Util "./util"
)

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
	server, err := net.Listen(Util.PROTOCOL, Util.PORT)
	var adressList, msgList []string

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

		go handleClient(client, &adressList, &msgList)
	}
}

// This function takes charge of handling clients
// by sending them a process the first time they
// connect to the server
func handleClient(client net.Conn, clientsList *[]string, msgList *[]string) {
	// wait for client ID
	data := make([]byte, 100)
	br, err := client.Read(data)

	// terminate when an error ocurrs
	if err != nil {
		fmt.Println(err)
		return
	} else {
		clientID := strings.Split(string(data[:br]), "|")[0]
		// check if client already exist
		if clientIsInList(clientID, *clientsList) {
			handleData(string(data[:br]), msgList)
		} else {
			fmt.Println(clientID + " connected...")
			*clientsList = append(*clientsList, clientID)
		}
	}
}

// Handles the data from the clients
// The data such as: message, files and exit
func handleData(data string, msgList *[]string) {
	var dataContent = strings.Split(data, Util.Separator)

	switch dataContent[1] {
	case Util.File:
		break
	case Util.Message:
		*msgList = append(*msgList, dataContent[0]+Util.Space+dataContent[2])
		fmt.Println(dataContent[0] + Util.Space + dataContent[2])
		break
	case Util.Exit:
		fmt.Println(dataContent[0] + " disconnected...")
		*msgList = Util.Remove(*msgList, dataContent[0])
		break
	default:
		break
	}
}

func main() {
	go server()

	// press enter to exit...
	fmt.Scanln()
}
