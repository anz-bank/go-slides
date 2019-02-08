// Defer
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Deferred function calls are executed in Last In First Out order
// after the surrounding function returns
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func main() {
	if err := ioutil.WriteFile("/tmp/hi.txt", []byte("hello_world\n"), 0644); err != nil {
		log.Fatal(err)
	}
	if _, err := CopyFile("/tmp/hi2.txt", "/tmp/hi.txt"); err != nil {
		log.Fatal(err)
	}
	txt, err := ioutil.ReadFile("/tmp/hi2.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(txt))
}
