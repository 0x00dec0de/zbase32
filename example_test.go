package zbase32_test

import (
	"bytes"
	"fmt"
	"os"

	"gopkg.in/corvus-ch/zbase32.v0"
)

func Example() {
	s := zbase32.EncodeToString([]byte{240, 191, 199})
	fmt.Println(s)
	// Output:
	// 6n9hq
}

func ExampleNewEncoder() {
	input := []byte("foo\x00bar")
	encoder := zbase32.NewEncoder(os.Stdout)
	encoder.Write(input)
	// Must close the encoder when finished to flush any partial blocks.
	encoder.Close()
	// Output:
	// c3zs6ydncf3y
}

func ExampleNewDecoder() {
	input := []byte("c3zs6ydncf3y")
	decoder := zbase32.NewDecoder(bytes.NewReader(input))
	output := make([]byte, 16)
	n, err := decoder.Read(output)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(output[:n])
	// Output:
	// [102 111 111 0 98 97 114]
}
