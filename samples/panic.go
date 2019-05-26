// panic
package main

import "os"

func main() {
	_, err := os.Create("/nonexistent/file")
	if err != nil {
		panic(err)
	}
}
