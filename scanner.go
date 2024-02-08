// Filename: Scanner.go
// Code for a scanner for MySQL. Code will be triggered from Main.go, and used to perform the scanning action.
// Language: Go
// To-do: MYSQL Handshake

package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

type arguments struct {
	Host string
	Port string // if no port specified, using port
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

func testPrint() {
	fmt.Println("Port Scanning")
	open := scanPort("tcp", ipaddr, port_num) // default port
	fmt.Printf("Port Open: %t\n", open)
}

func scanPort(protocol, hostname string, port int) bool {
	fmt.Println("Hit scanner.scanPort")
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		return false
	}

	defer conn.Close()
	return true
}
