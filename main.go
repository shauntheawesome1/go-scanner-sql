// Filename: Main.go
// Code for a scanner for MySQL. Code will be triggered from Main.go, and used to perform the scanning action.
// Language: Go
// To-do: MYSQL Handshake

package main

import (
	"fmt"
)

var ipaddr string
var port_num int

func main() {

	fmt.Println("Hit Main.main.")
	fmt.Println("Enter Your IP Address: ")

	fmt.Scanln(&ipaddr)
	if ipaddr == "" || ipaddr == "localhost" || ipaddr == "default" {
		ipaddr = "127.0.0.1"
	}

	fmt.Println("Enter Port # to Scan: ")

	fmt.Scanln(&port_num)

	if port_num < 1 || port_num > 65535 {
		port_num = 3306
	}
	fmt.Println(ipaddr)
	testPrint()

}
