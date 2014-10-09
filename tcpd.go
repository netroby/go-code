package main

import "net"

const (
	BIND_HOST = "0.0.0.0"
	BIND_PORT = "18080"
	BIND_TYPE = "tcp"
)

func main() {
	l, err := net.Listen(BIND_TYPE, BIND_HOST+":"+BIND_PORT)
	if err != nil {
		panic(err.Error())
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err.Error())
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		panic(err.Error())
	}
	conn.Write([]byte("Message received!"))
	conn.Close()
}
