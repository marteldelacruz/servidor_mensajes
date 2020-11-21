package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"

	Util "./util"
)

var showMessage = false
var endQueue = false

// The new client will connect to a server on the PORT
// if the server is not running already, an error message
// will be shown
func client(id string) {
	conn, err := net.Dial(Util.PROTOCOL, Util.PORT)

	if err != nil {
		fmt.Println(err)
		return
	}
	// send client ID
	conn.Write([]byte(id))

	go receiveData(id)
}

// This is a go routine that will manage all incomming
// data from the server
func receiveData(id string) {
	data := make([]byte, Util.Max_File_Size)
	conn, err := net.Dial(Util.PROTOCOL, Util.PORT)

	if err != nil {
		fmt.Println(err)
		return
	}

	// send client ID
	conn.Write([]byte(id + Util.Separator + Util.Ask))

	// loop to handle server messages
	for {
		br, err := conn.Read(data)

		if err != nil {
			fmt.Println(err)
			continue
		}

		dataSlice := strings.Split(string(data[:br]), Util.Separator)

		if showMessage {
			switch dataSlice[0] {
			// print message
			case Util.Messages:
				fmt.Println(dataSlice[1])
				break
			// save file
			case Util.File:
				Util.SaveFile(dataSlice[1], dataSlice[2])
				break
			}
		}
	}
}

// Sends a data type to the server.
// The data can be a message or a exit signal
// The id represents the client name, d_type the data type
// to be sent and m the message
func sendData(id string, d_type string, m string) {
	conn, err := net.Dial(Util.PROTOCOL, Util.PORT)

	if err != nil {
		fmt.Println(err)
		return
	}

	// send client ID and message with separator
	conn.Write([]byte(id + Util.Separator + d_type + Util.Separator + m))
}

// Sends a file to the server
func sendFile(id string) {
	var path string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("File path: ")
	scanner.Scan()
	path = scanner.Text()
	bytes := Util.GetFile(path)
	sendData(id, Util.File+Util.Separator+filepath.Base(path), bytes)
}

// The main menu contains all the available options
// for the client
func mainMenu(id string) {
	var opt, msg string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\n\n")
		fmt.Println("1-. Send message")
		fmt.Println("2.- Send file")
		fmt.Println("3.- View chat")
		fmt.Println("0.- Exit")

		fmt.Print("Select an option: ")
		scanner.Scan()
		opt = scanner.Text()

		switch opt {
		case "1":
			fmt.Print("Type message: ")
			scanner.Scan()
			msg = scanner.Text()
			sendData(id, Util.Message, msg)
			break
		case "2":
			sendFile(id)
			break
		case "3":
			showMessage = true
			scanner.Scan()
			scanner.Text()
			showMessage = false
			break
		case "0":
			sendData(id, Util.Exit, "")
			return
			break
		default:
		}
	}
}

func main() {
	var id string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Nickname: ")
	scanner.Scan()
	id = scanner.Text()

	client(id)

	mainMenu(id)
}
