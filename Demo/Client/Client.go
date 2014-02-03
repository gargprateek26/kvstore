package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"
)

func Client(port_num string) {
	var msg string
	port_num = ":"+port_num
	// connect to the server
	c, err := net.Dial("tcp", "127.0.0.1"+port_num)
	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(os.Stdin) //creates a buffer to read from terminal

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			// You may check here if err == io.EOF
			break
		}
		//fmt.Println("The line read is: " + string(line))
		fmt.Println("Sent:", line)
		err = gob.NewEncoder(c).Encode(line)
		if err != nil {
			fmt.Println(err)
		}
		err = gob.NewDecoder(c).Decode(&msg) //sends the message read from terminal
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Received:", msg)
		}

	}

	c.Close()
}

func main() {
	
	reader := bufio.NewReader(os.Stdin) //creates a buffer to read from terminal
	fmt.Println("Enter Port to connect to: ")
	port, _ := reader.ReadString('\n')
	port = strings.TrimSpace(port)
	
	Client(port)

	//var input string
	//fmt.Scanln(&input)
}
