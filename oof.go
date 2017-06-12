package main

import "fmt"

func main() {
	var l int
	var b int
	println("enter a big number.")
	_, err := fmt.Scanf("%d", &b)
	if err != nil {
		println("ERROR: ", err)
	}
	println("now enter a number smaller number than", b)
	_, err = fmt.Scanf("%d", &l)
	if err != nil {
		println("ERROR: ", err)
	}

	println(b % l)
}
