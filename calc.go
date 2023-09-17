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
	var err error
	var res float64
	if res, err = calculator.Calculate(expression); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
