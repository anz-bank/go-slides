package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2000 * time.Millisecond)
		fmt.Fprint(w, "pong")
	})
	addr := ":9090"
	fmt.Println("Starting Josh's webserver on port", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
