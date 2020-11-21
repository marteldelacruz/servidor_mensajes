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
		for {
			// client is out
			if clientList[i].Exit {
				break
			}
			if newMessage {
				clientList[i].Conn.Write([]byte(Util.Messages + Util.Separator + Util.ListToString(*msgList)))
			}
			time.Sleep(time.Millisecond * 500)
		}
	}
}

// sends a file to all the clients
func sendFileToClients(clientList []Util.Client, name string, content string) {
	for _, client := range clientList {
		client.Conn.Write([]byte(Util.File + Util.Separator + name + Util.Separator + content))
	}
}

// Saves a new client on the client list
func saveNewClient(id string, clientList *[]Util.Client) {
	fmt.Println(id + " CONNECTED...")
	newClient := Util.Client{
		Name: id,
		Conn: nil,
		Exit: false,
	}
	*clientList = append(*clientList, newClient)
}

// Updates the client connection obj on the list
func updateClient(id string, client net.Conn, clientList *[]Util.Client, msgList *[]string) {
	i := Util.GetClientIndex(id, *clientList)

	// client exist
	if i != Util.Invalid {
		(*clientList)[i].Conn = client
	}

	// send data to client
	go sendDataToClient(id, *clientList, msgList)
}

// Removes a client from the client list using the id as reference
func removeClient(id string, clientList *[]Util.Client) {
	i := Util.GetClientIndex(id, *clientList)

	// client exist
	if i != Util.Invalid {
		*clientList = Util.RemoveIndex(*clientList, i)
	}
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
		dataSlice := strings.Split(string(data[:br]), Util.Separator)

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
	id := dataContent[0]

	switch dataContent[1] {
	case Util.Ask: // Ask for data
		fmt.Println(id + " ASK FOR DATA")
		updateClient(id, client, clientList, msgList)
		break
	case Util.Exit: // Exit from the server
		fmt.Println(id + " DISCONNECTED")
		removeClient(id, clientList)
		break
	case Util.File: // Receive a file
		fmt.Println(id + " SENT A FILE")
		sendFileToClients(*clientList, dataContent[2], dataContent[3])
		break
	case Util.Message: // Receive a message
		fmt.Println(id + " SENT A MESSAGE")
		newMessage = true
		*msgList = append(*msgList, id+Util.Space+dataContent[2])
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
