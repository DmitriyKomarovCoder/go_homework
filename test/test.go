package main

import (
	"fmt"
	"strings"
)

func main() {
	outString := []string{"12221", "sdfsd", "fdsf"}
	str := strings.Join(outString[1:], " ")
	fmt.Println(str)
}
