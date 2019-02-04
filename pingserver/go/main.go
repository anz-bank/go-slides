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
	fmt.Println("Starting webserver on port 9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
