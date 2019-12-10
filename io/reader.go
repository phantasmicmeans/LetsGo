package main

/*
Readers

The io package specifies the io.Reader interface,
which represents the read end of a stream of data

The Go standard library contains many implementations of there interfaes,
including files, network connections, compressors, ciphers, and others/

The io.Reader interface has a Read method:
	func (T) Read (b []byte) (n int, err error)

Read populates the given byte slice with data
and returns the number of bytes populated
and an error value.
It returns an io.EOF error when the stream ends.

The examples code creates a strings.Reader
and consumes its output 8 bytes at a time.
*/
import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v\n", n, err)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
