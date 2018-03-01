package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	SocketServer()
}

func SocketServer() {
	listen, err := net.Listen("tcp", ":8888")
	defer listen.Close()
	failOnError("Socket failt to listen port: %s", err)

	log.Printf("Server started")

	for {
		conn, err := listen.Accept()
		failOnError("Failed to accept connection", err)

		go handler(conn)
	}
}

func handler(conn net.Conn) {
	defer conn.Close()

	var (
		buf = make([]byte, 1024)
		r   = bufio.NewReader(conn)
		// w   = bufio.NewWriter(conn)
	)
ILOOP:
	for {
		n, err := r.Read(buf)

		switch err {
		case io.EOF:
			break ILOOP
		case nil:
			messageReceived := string(buf[:n])
			log.Println("Message Received: ", messageReceived)
			if strings.HasSuffix(messageReceived, "\r\n\r\n") {
				break ILOOP
			}
		default:
			log.Fatalf("Failed to receive data: %s", err)
			return
		}

		// w.Write([]byte("Tks for your message \n"))
		// w.Flush()
		conn.Write([]byte("Tks for your message \n"))
	}
}

func failOnError(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
