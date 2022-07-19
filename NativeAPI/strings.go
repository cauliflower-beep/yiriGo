package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"111", "222", "333", "444", "555"}
	str := strings.Join(s, ",")
	fmt.Println(str)
}
