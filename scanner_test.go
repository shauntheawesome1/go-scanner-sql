package main

import (
	"strconv"
	"testing"
	"time"

	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/mysql"
	"github.com/stretchr/testify/assert"
)

type sql_handshake struct {
	version_name    string
	plugin_name   string
	plugin_data    string
	flag_status    string
	char_set       string
	conn_id        string
	handshake_type int
}

func init(){
	fmt.Println("Hit scanner_test.go")
}

func TestConnect(protocol, hostname string, port int) {
	fmt.Println("Hit scannet_test.go - TestConnect()")
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)
}