package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "   abcd"
	b := "abcd   "
	c := "  abc  "
	d := " a b c "
	fmt.Println(strings.TrimSpace(a),
		strings.TrimSpace(b),
		strings.TrimSpace(c),
		strings.TrimSpace(d))
}
