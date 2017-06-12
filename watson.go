package main

import "fmt"

func main() {
	a := fmt.Println
	for i := 0; i < 101; i++ {
		a(i)
	}
}
