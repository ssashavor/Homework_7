package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func handleConnection(conn net.Conn) {
	var response string
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err == nil {
			response = strconv.Itoa(num * 2)
		} else {
			response = strings.ToUpper(scanner.Text())
		}
		conn.Write([]byte(response))
	}
	if err := scanner.Err(); err != nil {
		fmt.Errorf("Scanning error", err)
	}

}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		go handleConnection(conn)
	}
}
