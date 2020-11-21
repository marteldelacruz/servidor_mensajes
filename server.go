package main

import (
	"fmt"
	"net"
	"strings"
	"time"

	Util "./util"
)

var newMessage bool

/// This rutine runs the server on a loop to keep
/// handling client petitions using the TCP connection
/// on the 9999 port
func server() {
	server, err := net.Listen(Util.PROTOCOL, Util.PORT)
	var msgList []string
	var clientList []Util.Client

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
		go handleClient(client, &clientList, &msgList)
	}
}

// Keepds sending data to the client when a message is received
func sendDataToClient(id string, clientList []Util.Client, msgList *[]string) {
	i := Util.GetClientIndex(id, clientList)

	// client exist
	if i != Util.Invalid {
		fmt.Println(id + " exist on list")
		for {
			// client is out
			if clientList[i].Exit {
				break
			}
			if newMessage {
				fmt.Println("Msg sent to : " + id)
				clientList[i].Conn.Write([]byte(Util.ListToString(*msgList)))
			}
			time.Sleep(time.Millisecond * 500)
		}
	}
}

// Saves a new client on the client list
func saveNewClient(id string, clientList *[]Util.Client) {
	fmt.Println(id + " connected...")
	newClient := Util.Client{
		Name: id,
		Conn: nil,
		Exit: false,
	}
	*clientList = append(*clientList, newClient)
}

//
func updateClient(id string, client net.Conn, clientList *[]Util.Client, msgList *[]string) {
	i := Util.GetClientIndex(id, *clientList)

	// client exist
	if i != Util.Invalid {
		(*clientList)[i].Conn = client
	}

	// send data to client
	go sendDataToClient(id, *clientList, msgList)
}

// This function takes charge of handling clients
// by sending them a process the first time they
// connect to the server
func handleClient(client net.Conn, clientList *[]Util.Client, msgList *[]string) {
	// wait for client ID
	data := make([]byte, Util.Max_File_Size)
	br, err := client.Read(data)

	// terminate when an error ocurrs
	if err != nil {
		fmt.Println(err)
		return
	} else {
		dataSlice := strings.Split(string(data[:br]), "|")

		// check if is client already exist
		if Util.IsInList(dataSlice[0], *clientList) {
			handleData(client, clientList, string(data[:br]), msgList)
		} else { // save new client
			saveNewClient(dataSlice[0], clientList)
		}
	}
}

// Handles the data from the clients
// The data such as: message, files and exit
func handleData(client net.Conn, clientList *[]Util.Client, data string, msgList *[]string) {
	var dataContent = strings.Split(data, Util.Separator)

	switch dataContent[1] {
	case Util.Ask: // Ask for data
		fmt.Println(dataContent[0] + " ask for data")
		updateClient(dataContent[0], client, clientList, msgList)
		break
	case Util.Exit: // Exit from the server
		fmt.Println(dataContent[0] + " disconnected...")
		*msgList = Util.Remove(*msgList, dataContent[0])
		break
	case Util.File: // Receive a file
		fmt.Println(dataContent[0] + " sent a file...")
		fmt.Println("bytes => " + dataContent[2])
		break
	case Util.Message: // Receive a message
		newMessage = true
		*msgList = append(*msgList, dataContent[0]+Util.Space+dataContent[2])
		fmt.Println(dataContent[0] + Util.Space + dataContent[2])
		time.Sleep(time.Millisecond * 500)
		newMessage = false
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
