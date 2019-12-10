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
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		conn.Write([]byte(multipliedInput(scanner.Text())))
	}
	if err := scanner.Err(); err != nil {
		fmt.Errorf("Scanning error", scanner.Err())
	}

}

func multipliedInput(input string) string {
	num, err := strconv.Atoi(input)
	if err == nil {
		return strconv.Itoa(num * 2)
	}
	return strings.ToUpper(input)
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