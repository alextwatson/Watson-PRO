package main

import "fmt"

func main() {
	for i := -321; i < 321; i++ {
		fmt.Printf("%b \t %d \n", i, i)
	}
}
