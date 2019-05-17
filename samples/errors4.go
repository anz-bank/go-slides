package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("https://example.org:^/?uid=john&prj=oss")
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	fmt.Println(q.Get("uid"))
	fmt.Println(q.Get("prj"))
}
