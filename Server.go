package kvstore

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

var logFile *os.File
var data_store map[string]string = make(map[string]string) //creates a hash map
var mutex = &sync.Mutex{}

func Server(port_num string) {
	// listen on a port
	port_num = ":" + port_num
	ln, err := net.Listen("tcp", port_num)
	if err != nil {
		fmt.Println(err)
		return
	}

	//log.Println("Started Server on port number", port_num)
	acceptConn(ln)
}

func acceptConn(ln net.Listener) {
	//log.Println("Server Accepting Connections")
	for {

		// accept a connection
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// handle the connection
		go handleServerConnection(c)
	}
}

func Get_data(key string) (string, bool) { //function to retrieve a value for key in hash map

	var str string
	var flag bool
	mutex.Lock()
	str, flag = data_store[key]
	mutex.Unlock()
	return str, flag
}

func Set_data(key, val string) bool { //function to set value for a key in hash map

	mutex.Lock()
	data_store[key] = val
	mutex.Unlock()
	return (true)
}

func Del_data(key string) bool { //function to delete an existing key in hash map
	mutex.Lock()
	delete(data_store, key)
	mutex.Unlock()
	return (true)
}

func handleServerConnection(c net.Conn) {
	// receive the message
	for {
		var msg, action, key, value, val_ret string
		var success_flag bool
		var parts []string
		err := gob.NewDecoder(c).Decode(&msg)
		if err != nil {
			fmt.Println(err)
		} else {
			//log.Println("Received: ", msg)
		}
		parts = strings.Split(msg, " ")

		action = parts[0] //breaks the input to command, key, value
		key = parts[1]    //retrieves the key from input
		switch action {
		case "get":
			key = strings.TrimSpace(key) //trims the new line character
			val_ret, success_flag = Get_data(key)
			key = strings.TrimSpace(val_ret)
			if success_flag {
				err = gob.NewEncoder(c).Encode(val_ret)
				//log.Println("Value of " + val_ret + " returned")
				if err != nil {
					log.Println(err)
				}
			} else {
				err = gob.NewEncoder(c).Encode("Error")
				//log.Println("Error in Get")
				if err != nil {
					log.Println(err)
				}
			}

		case "set":
			value = parts[2] //retrieves the value from input
			success_flag = Set_data(key, value)
			val_ret, success_flag = Get_data(key)
			if val_ret == value && success_flag == true { //fmt.Println(get_data(key))
				//err = gob.NewEncoder(c).Encode("Success")
				//log.Println("Value of " + key + " set to " + value)
				if err != nil {
					log.Println(err)
				}
			} else {
				err = gob.NewEncoder(c).Encode("Error")
				log.Println("Error in setting Value")
				if err != nil {
					log.Println(err)
				}
			}

		case "del":
			key = strings.TrimSpace(key) //retrieves the newline character from key
			val_ret, success_flag = data_store[key]
			if success_flag == true {
				success_flag = Del_data(key)
				//err = gob.NewEncoder(c).Encode("Success")
				//log.Println("Key " + key + " deleted successfully")
				if err != nil {
					log.Println(err)
				}
			} else {
				err = gob.NewEncoder(c).Encode("Error")
				//log.Println("Key " + key + " does not exist")
				if err != nil {
					log.Println(err)
				}
			}

		default:
			err = gob.NewEncoder(c).Encode("Error")
			//log.Println("Command Error")
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	c.Close()
}

/*
func main() {
	var err error
	logFile, err = os.Create("logfile.txt")
	if err != nil {
		log.Fatal("Log file create:", err)
		return
	}
	log.SetOutput(logFile)

	reader := bufio.NewReader(os.Stdin) //creates a buffer to read from terminal
	//fmt.Println("Enter Port to connect to: ")
	port, err := reader.ReadString('\n')
	port = strings.TrimSpace(port)

	Server(port) //initializes the server with the given port number
	//data_store = make(map[string]string)
	//var input string
	//fmt.Scanln(&input)
}
*/
