package util

import (
	"fmt"
	"io/ioutil"
	"net"
	"os/user"
)

const (
	Ask           = "0"
	Exit          = "4"
	File          = "2"
	Invalid       = -1
	Max_File_Size = 1000
	Message       = "1"
	Messages      = "3"
	PROTOCOL      = "tcp"
	PORT          = ":9999"
	Separator     = "|"
	Space         = ": "
)

type Client struct {
	Name string
	Conn net.Conn
	Exit bool
}

// Returns the given string list as a single string
func ListToString(list []string) string {
	s := "\n---------------------------------------\n"

	for _, b := range list {
		s += b + "\n"
	}

	return s + "---------------------------------------"
}

// Verifies if a string already exist on the list
// and then returns its index on the list
func GetClientIndex(a string, list []Client) int {
	for i, b := range list {
		if b.Name == a {
			return i
		}
	}
	return Invalid
}

// Verifies if a string already exist on the list
func IsInList(a string, list []Client) bool {
	for _, b := range list {
		if b.Name == a {
			return true
		}
	}
	return false
}

// Removes an element from the string list
func Remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// Saves a new file with the given name and content
func SaveFile(name string, content string) {
	myself, error := user.Current()

	if error != nil {
		fmt.Println(error)
	}

	fullPath := myself.HomeDir + "/Documents/" + name
	fmt.Println("FILE SAVED AT => " + fullPath)
	ioutil.WriteFile(fullPath, []byte(content), 0)
}

// Returns a string with the bytes of the file
func GetFile(path string) string {
	b, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Print(err)
		return ""
	}

	return string(b)
}
