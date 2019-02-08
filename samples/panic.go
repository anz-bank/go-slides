// Panic
package main

import "os"

func main() {
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
