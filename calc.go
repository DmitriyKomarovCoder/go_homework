package main

import (
	"calculator/calculator"
	"fmt"
	"os"
	"strings"
)

func main() {
	arg := os.Args[1:]
	expression := strings.Join(arg, " ")
	res, err := calculator.Calculate(expression)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
