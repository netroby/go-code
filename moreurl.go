package main

import (
	"io"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, World!\n")
}

func DudeServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Dude, what's going on\n")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/", DudeServer)
	err := http.ListenAndServe(":18080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
