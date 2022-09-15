package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func echo(conn net.Conn) {
	r := bufio.NewReader(conn)
	for {
		line, err := r.ReadBytes(byte('\n'))
		switch err {
		case nil:
			break
		case io.EOF:
		default:
			fmt.Println("ERROR", err)
		}

		var remoteIP = ""
		if addr, ok := conn.RemoteAddr().(*net.TCPAddr); ok {
			remoteIP = addr.IP.String()
		}

		conn.Write([]byte(remoteIP + ": " + string(line[:]) + "\n"))
	}
}

func main() {
	var hostport = os.Getenv("HOSTPORT")

	l, err := net.Listen("tcp", hostport)
	if err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()

		if err != nil {
			fmt.Println("ERROR", err)
			continue
		}
		go echo(conn)
	}
}
