package main

import (
	"bytes"

	"golang.org/x/net/html"
)

func main() {
	_, _ = html.Parse(bytes.NewReader([]byte("test")))
}
