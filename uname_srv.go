package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"net"
)

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
	_, err := conn.Read(buf[:])
	if err != nil {
		panic(err.Error())
	}
	var uid string
	uid = string(buf)
	var name string
	name = getName(uid)
	conn.Write([]byte(name))
	conn.Close()
}
func getName(uid string) string {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/x_ucenter?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	stmtOut, err := db.Prepare("Select username from uc_members where uid = ? limit 1;")
	if err != nil {
		panic(err.Error())
	}
	var username string
	err = stmtOut.QueryRow(uid).Scan(&username)
	if err != nil {
		panic(err.Error())
	}
	return username
}
