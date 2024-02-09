// Filename: Scanner.go
// Code for a scanner for MySQL. Code will be triggered from Main.go, and used to perform the scanning action.
// Language: Go
// To-do: MYSQL Handshake

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"strconv"
	"time"
)

type ScanOutput struct {
	Host    string
	Port    int    // if no port specified, using port
	State   string // state of specified port
	Message int    // connection message
}

type sql_evidence struct {
	version_set    string
	plugin_name    string
	plugin_data    string
	flag_status    string
	char_set       string
	conn_id        string
	handshake_type int
}

func init() {
	fmt.Println("Hit Scanner.init")

}

/*func testPrint() {
	fmt.Println("Port Scanning")

	open := scanPort("tcp", ipaddr, port_num) // default port = 1332
	fmt.Printf("Port Open: %t\n", open)
}*/

func scanPort(protocol, hostname string, port int) ScanOutput {
	fmt.Println("Hit scanner.scanPort")
	output := ScanOutput{Host: hostname, Port: port}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		output.State = "Closed"
		return output
	}
	fmt.Printf("Port: %d is open\n\n", port)

	if err != nil {
		output.State = "Closed"
		return output
	}

	//defer conn.Close()
	/*fmt.Println("Connection Closed")
	buf := make([]byte, 4)
	fmt.Println("Created Buffer")
	output.Message, err = conn.Read(buf)
	fmt.Println("Done Reading Connection")
	output.State = "Open"*/

	/*lenBuf := make([]byte, 4)
	fmt.Println("Created Buffer")
	output.Message, err = conn.Read(lenBuf)
	fmt.Println("Done Reading Connection")
	if err != nil {
		output.State = "Closed"
		return output
	}*/

	lenBuf := make([]byte, 4)
	fmt.Println("Created Buffer1")
	_, err = conn.Read(lenBuf)
	fmt.Println("Connection Passed1")
	if err != nil {
		output.State = "Closed"
		return output
	}

	lenData, err := msgLength(lenBuf)
	if err != nil {
		output.State = "Closed"
		return output
	}

	/*buf := make([]byte, lenData)
	output.Message, err = conn.Read(buf)
	if err == io.EOF {
		output.State = "Closed"
		return output
	}
	if err != nil {
		output.State = "Closed"
		return output
	}*/

	buf := make([]byte, lenData)
	fmt.Println("Created Buffer2")
	_, err = conn.Read(buf)
	fmt.Println("Connection Passed2")
	if err == io.EOF {
		output.State = "Closed"
		return output
	}
	if err != nil {
		output.State = "Closed"
		return output
	}

	if err := conn.Close(); err != nil {
		output.State = "Closed"
		return output
	}

	return output
}

func InitialScan(hostname string, port_num int) []ScanOutput {

	var results []ScanOutput

	fmt.Println("Port Scanning")

	results = append(results, scanPort("tcp", ipaddr, port_num)) // default port = 1332
	fmt.Printf("Port Open: %d\n", port_num)

	return results
}

func msgLength(b []byte) (int32, error) {
	buf := bytes.NewReader(b)
	var result int32
	err := binary.Read(buf, binary.LittleEndian, &result)

	return result, err
}
