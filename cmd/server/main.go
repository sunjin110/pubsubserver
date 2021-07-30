package main

import (
	"fmt"
	"io"
	"net"
	"pubsub/pkg/common/chk"
)

func main() {
	fmt.Println("tcp pub sub server")
	server()
}

func server() {

	listener, err := net.Listen("tcp", "localhost:10000")
	chk.SE(err)

	fmt.Println("Server running at localhost:10000")
	waitClient(listener)
}

func waitClient(listener net.Listener) {

	for {
		connection, err := listener.Accept()
		chk.SE(err)
		go goEcho(connection) // 別threadでやりとりする
	}

}

func goEcho(connection net.Conn) {
	defer connection.Close()

	for {
		var buf = make([]byte, 1024)

		n, err := connection.Read(buf)
		if err != nil {
			if err == io.EOF {
				return
			}
			chk.SE(err)
		}

		fmt.Printf("Client> %s\n", string(buf[:n]))

		_, err = connection.Write(buf[:n])
		chk.SE(err)
	}
}
