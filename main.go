package main

import (
	"bytes"
	"fmt"

	"github.com/tobocop/go-things/json"
)

func main() {
	cw := &capturingWriter{nil}

	json.NewEncoder(cw).Encode("some-initial-data")

	json.NewEncoder(&bytes.Buffer{}).Encode("nope")
	fmt.Println(string(cw.data))
}

type capturingWriter struct {
	data []byte
}

func (cw *capturingWriter) Write(p []byte) (int, error) {
	cw.data = p

	return len(p), nil
}
