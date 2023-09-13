package main

import (
	"fmt"
	"io"
	"strings"
)


func DemoStringReader() {
    r := strings.NewReader("Hello, Reader!")

	buffer := make([]byte, 8) // allocate buffer to consume reader output 8 bytes at a time
	for {
        // n : number of bytes successfully populated in buffer
		n, err := r.Read(buffer)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, buffer)
		fmt.Printf("b[:n] = %q\n", buffer[:n])
		if err == io.EOF {
			break
		}
	}
}

// a read method is implemented by the below type; it means we can call read method from Reader interface
// in this case we simply emit a stream of the char 'A'
type MyAReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (mr MyAReader) Read(b []byte) (n int, err error) {
	//blen := len(b)
	for i := range b {
		b[i] = 'A' // in other readers like stringReader, the buffer is to be populated with strings(bcoz stringReader). here in this oversimplified example, we just emit a stream of char 'A', and read from nothing in particular.
		// In other situations, we could be reading from say a file and emit a stream of file contents
	}

	// infinite stream means no return of io.EOF error
	return len(b), nil
}

