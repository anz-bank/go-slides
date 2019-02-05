package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	addr := ":4242"
	fmt.Printf("Open your web browser and visit http://127.0.0.1%s\n", addr)
	err := http.ListenAndServe(addr, http.FileServer(http.Dir(".")))
	if err != nil {
		log.Printf("Error running web server for static assets: %v", err)
	}
}
