package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	host := os.Args[1]
	dnsServer(host)
}

func dnsServer(host string) {
	ln, err := net.Listen("tcp", ":5353")
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			fmt.Println(err)
		}
		go handleConnection(conn, host)
	}
}

func handleConnection(conn net.Conn, host string) {
	// fmt.Println("lookup host:", host)
	// ip, err := net.LookupHost("google.com")
	ip, err := net.LookupIP(host)
	getCloudFlare(conn, host)
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	fmt.Println(ip)
	conn.Close()
}

func getCloudFlare(conn net.Conn, host string) {
	conn, err := net.Dial("tcp", "1.1.1.1:853")
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	fmt.Println("inside getCloudFlare", host)
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("status:", status)
}
