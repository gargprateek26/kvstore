package kvstore

import (
	"encoding/gob"
	"fmt"
	"net"
	"strconv"
	"testing"
	"time"
)

//var test_data_store map[string]string = make(map[string]string) //creates a hash map
var conn_map map[int]net.Conn = make(map[int]net.Conn) //creates the hash map to store connection info

func TestSetSingle(t *testing.T) { //this function tests for a single SET and single GET by different clients

	var conn_var net.Conn
	var msg string

	go Server("9999") //initialises the server on Port 9999

	time.Sleep(10 * time.Millisecond)

	for j := 0; j < 2; j++ { //creates two clients
		conn_var = newClient()
		conn_map[j] = conn_var
	}

	line := "set hello 123"
	err := gob.NewEncoder(net.Conn(conn_map[0])).Encode(line)
	if err != nil {
		t.Errorf("Error Found")
	}
	time.Sleep(1 * time.Millisecond)

	line = "get hello"
	err = gob.NewEncoder(net.Conn(conn_map[1])).Encode(line)
	if err != nil {
		t.Errorf("Error Found")
	}
	err = gob.NewDecoder(conn_map[1]).Decode(&msg) //sends the message read from terminal
	if err != nil {
		t.Errorf("Error Found")
	}

	if msg != "123" {
		t.Errorf("Error in retrieval", msg)
	}
	fmt.Println("Single set testing done")

}

func TestMulti(t *testing.T) { //this function tests for a SET and GET by different clients

	var conn_var net.Conn
	var msg, key, value string
	numClients := 4
	numIter := 4
	//go Server("9999") //initializes the server on Port 9999

	time.Sleep(10 * time.Millisecond)

	//*************************tests for several SETs followed by GET****************************************
	for j := 0; j < numClients; j++ { //creates numClients clients
		conn_var = newClient()

		conn_map[j] = conn_var
		for i := 1; i <= numIter; i++ {
			key = "C" + strconv.Itoa((j*10)+i)
			value = "c" + strconv.Itoa((j*10)+i)
			line := "set " + key + " " + value
			err := gob.NewEncoder(net.Conn(conn_map[j])).Encode(line)
			if err != nil {
				t.Errorf("Error Found")
			}

			time.Sleep(10 * time.Millisecond)
		}
	}

	line := "get C12"

	err := gob.NewEncoder(net.Conn(conn_map[1])).Encode(line)
	if err != nil {
		t.Errorf("Error Found")
	}

	err = gob.NewDecoder(conn_map[1]).Decode(&msg) //sends the message read from terminal
	if err != nil {
		t.Errorf("Error Found")
	}
	if msg != "c12" {
		t.Errorf("Error in retrieval", msg)
	}
	fmt.Println("Several SET followed by GET tested")
	//***************************************************************************************************************
	//#########################tests for SET followed by SET followed by GET################################################
	key = "C13"
	value = "c11"
	line = "set " + key + " " + value
	err = gob.NewEncoder(net.Conn(conn_map[2])).Encode(line)
	if err != nil {
		t.Errorf("Error Found")
	}

	time.Sleep(1 * time.Millisecond)

	key = "C13"
	value = "c10"
	line = "set " + key + " " + value
	err = gob.NewEncoder(net.Conn(conn_map[2])).Encode(line)
	if err != nil {
		t.Errorf("Error Found")
	}

	time.Sleep(1 * time.Millisecond)

	line = "get C13"
	err = gob.NewEncoder(net.Conn(conn_map[1])).Encode(line)
	if err != nil {
		t.Errorf("Error Found")
	}
	err = gob.NewDecoder(conn_map[1]).Decode(&msg) //sends the message read from terminal
	if err != nil {
		t.Errorf("Error Found")
	}
	if msg != "c10" {
		t.Errorf("Error in retrieval", msg)
	}
	fmt.Println("SET SET GET Tested")
	//#######################################################################################################################
	//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^test for SET followed by DEL followed by GET^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

	key = "C14"
	value = "c12"
	line = "set " + key + " " + value
	err = gob.NewEncoder(net.Conn(conn_map[2])).Encode(line)
	if err != nil {
		t.Errorf("Error Found")
	}
	time.Sleep(2 * time.Millisecond)

	line = "del C14"
	err = gob.NewEncoder(net.Conn(conn_map[1])).Encode(line)
	if err != nil {
		t.Errorf("Error Found")
	}

	time.Sleep(1 * time.Millisecond)

	line = "get C14"
	err = gob.NewEncoder(net.Conn(conn_map[1])).Encode(line)
	if err != nil {
		t.Errorf("Error Found")
	}

	err = gob.NewDecoder(conn_map[1]).Decode(&msg) //sends the message read from terminal
	if err != nil {
		t.Errorf("Error Found")
	}
	if msg != "Error" {
		t.Errorf("Error in retrieval", msg)
	}
	fmt.Println("SET DEL GET Tested")
	//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

}

func newClient() net.Conn {

	c, err := net.Dial("tcp", "127.0.0.1:9999")

	if err != nil {
		fmt.Println(err)
	}
	return c
}
