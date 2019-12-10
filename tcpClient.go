package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	for {
		fmt.Print("Enter your message: ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		if text == "\n" {
			continue
		}
		if text == "exit\n" {
			return
		}
		if _, err := conn.Write([]byte(text)); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("Response: ")
		buff := make([]byte, 1024)
		_, err := conn.Read(buff)
		if err != nil {
			fmt.Errorf("Response error")
		}
		fmt.Println(string(buff))
	}
}

//Enter your message: hello
//Response: HELLO
//Enter your message: 123
//Response: 246
//Enter your message: exit
//(base) MacBook-Air-Sasha:server sashavorobyova$
