package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"pubsub/pkg/common/chk"
)

func main() {
	fmt.Println("tcp pub sub server")
	clinet()
}

func clinet() {

	connection, err := net.Dial("tcp", "localhost:10000")
	chk.SE(err)

	defer connection.Close()
	sendMessage(connection)
}

func sendMessage(connection net.Conn) {

	for {
		fmt.Print("> ")

		stdin := bufio.NewScanner(os.Stdin)
		if !stdin.Scan() {
			fmt.Println("good bye")
			return
		}

		_, err := connection.Write([]byte(stdin.Text()))
		chk.SE(err)

		var res = make([]byte, 4*1024)
		_, err = connection.Read(res)
		chk.SE(err)

		fmt.Printf("server> %s\n", res)
	}
}
