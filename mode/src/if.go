package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	var con = 'c'
	if con == 'c' {
		fmt.Println("condition is c")
	} else if con == 'a' {
		fmt.Println("condition is a")
	} else {
		fmt.Println("condition no found")
	}

	var x = 100
	if v := x - 101; v < 0 {
		fmt.Println("condition is v:", v)
	}
}