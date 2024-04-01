package main

import (
	"fmt"
	"io"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(slice []byte) (int, error) {
	readSlice := make([]byte, len(slice), cap(slice))

	readBytes, err := r.r.Read(readSlice)

	if err == nil {
		for i := 0; i < len(readSlice); i++ {
			if readSlice[i] <= 77 || (readSlice[i] >= 97 && readSlice[i] <= 109) {
				slice[i] = readSlice[i] + 13
			} else {
				slice[i] = readSlice[i] - 13
			}
		}
	}


	fmt.Println(readSlice)

	return readBytes, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}

	result := make([]byte, 50)

	n, _ := r.Read(result)

	fmt.Printf("%q", result[:n])

	// io.Copy(os.Stdout, &r)
}
