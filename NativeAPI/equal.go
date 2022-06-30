package main

import (
	"bytes"
	"fmt"
)

func main() {
	a := []byte{'a', 'b', 'c', 'd'}
	b := []byte{'a', 'b', 'c', 'd'}
	c := []byte{'a', 'b', 'c', 'd', 'e'}
	fmt.Println(bytes.Equal(a, b))
	fmt.Println(bytes.Equal(a, c))
}
