package main

import "fmt"

type Buffer struct {
	buf     []byte
	size    int
	maxSize int
}

func initBuf(maxSz int) *Buffer {
	buffer := &Buffer{
		buf:     make([]byte, 0),
		size:    0,
		maxSize: maxSz,
	}
	fmt.Printf("initBuf: size (%d), max(%d)\n", buffer.size, buffer.maxSize)
	return buffer
}

func (b *Buffer) Write(data []byte) error {
	if b.size+len(data) >= b.maxSize {
		return fmt.Errorf("buffer has reached max capacity")
	}
	b.buf = append(b.buf, data...)
	b.size += len(data)
	return nil
}

// Append that does too many checks
func (buf *Buffer) Append() {
	p0 := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	p1 := []byte{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	p2 := []byte{21, 22, 23, 24, 25, 26, 27, 28, 29, 30}

	err := buf.Write(p0[0:9])
	if err != nil {
		fmt.Println("Append: Error ", err)
		return
	}
	err = buf.Write(p1[0:9])
	if err != nil {
		fmt.Println("Append: Error ", err)
		return
	}
	err = buf.Write(p2[0:9])
	if err != nil {
		fmt.Println("Append: Error ", err)
		return
	}
	fmt.Println("Append: OK")
}

// Append that uses closure patter to reduce the
// number of checks.
func (buf *Buffer) AppendClosure() {
	var err error

	p0 := []byte{31, 32, 33, 34, 35, 36, 37, 38, 39, 40}
	p1 := []byte{41, 42, 43, 44, 45, 46, 47, 48, 49, 50}
	p2 := []byte{51, 52, 53, 54, 55, 56, 57, 58, 59, 60}

	write := func(data []byte) {
		if err != nil {
			return
		}
		err = buf.Write(data)
	}
	write(p0[0:9])
	write(p1[0:9])
	write(p2[0:9])
	if err != nil {
		fmt.Println("AppendClosure:", err)
		return
	}
	fmt.Println("AppendClosure: OK")
}

func main() {

	buf := initBuf(50)
	buf.Append()
	fmt.Println("Buffer after Append: ", buf.buf)
	buf.AppendClosure()
	fmt.Println("Buffer after AppendClosure: ", buf.buf)
}
